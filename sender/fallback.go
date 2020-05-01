/*
 * skogul, Fallback and dupe-sender
 *
 * Copyright (c) 2019 Telenor Norge AS
 * Author(s):
 *  - Kristian Lyngstøl <kly@kly.no>
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA
 * 02110-1301  USA
 */

package sender

import (
	"github.com/telenornms/skogul"
)

/*
Fallback sender tries each provided sender in turn before failing.

E.g.:

	primary := sender.InfluxDB{....}
	secondary := sender.Queue{....} // Not implemented yet
	emergency := sender.Debug{}

	fallback := sender.Fallback{}
	fallback.Add(&primary)
	fallback.Add(&secondary)
	fallback.Add(&emergency)

This will send data to Influx normally. If Influx fails, it will send it to
a queue. If that fails, it will print it to stdout.
*/
type Fallback struct {
	Next []*skogul.SenderRef `doc:"Ordered list of senders that will potentially receive metrics."`
}

// Send sends data down stream
func (fb *Fallback) Send(c *skogul.Container) error {
	for _, s := range fb.Next {
		err := s.S.Send(c)
		if err == nil {
			return nil
		}
	}
	return skogul.Error{Reason: "No working senders left..."}
}

// Dupe sender executes all provided senders in turn.
type Dupe struct {
	Next []*skogul.SenderRef `doc:"List of senders that will receive metrics, in order."`
}

// Send sends data down stream
func (dp *Dupe) Send(c *skogul.Container) error {
	var e error
	for _, s := range dp.Next {
		err := s.S.Send(c)
		if err != nil && e == nil {
			e = err
		}
	}
	return e
}

// logLog logs to a log. Loggingly.
var logLog = skogul.Logger("sender", "log")

/*
Log sender simply executes log.Print() on a predefined message.

Intended use is in combination with other senders, e.g. to explain WHY
sender.Debug() was used.
*/
type Log struct {
	Message string `doc:"Message to print."`
}

// Send logs a message and does no further processing
func (lg Log) Send(c *skogul.Container) error {
	logLog.Debug(lg.Message)
	return nil
}

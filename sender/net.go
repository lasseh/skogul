/*
 * skogul, net line-sender
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
	"encoding/json"
	"fmt"
	"net"

	"github.com/telenornms/skogul"
)

var netLog = skogul.Logger("sender", "net")

// Net sends metrics to a network address
type Net struct {
	Address string `doc:"Address to send data to" example:"192.168.1.99:1234"`
	Network string `doc:"Network, according to net.Dial. Typically udp or tcp."`
}

// Send sends metrics to a network address, json-encoded
func (n *Net) Send(c *skogul.Container) error {
	d, err := net.Dial(n.Network, n.Address)
	if err != nil {
		netLog.WithError(err).WithField("address", n.Address).Error("Failed to connect to target")
		return skogul.Error{Source: "net sender", Reason: "unable to connect to network address", Next: err}
	}
	// should almost certainly fix some method of retaining the
	// connection in the future
	defer d.Close()

	b, err := json.Marshal(c)
	if err != nil {
		return skogul.Error{Source: "net sender", Reason: "unable to marshal json for sending", Next: err}
	}
	nbytes, err := d.Write(b)
	if err != nil {
		return skogul.Error{Source: "net sender", Reason: "unable to send (all) data", Next: err}
	}
	if nbytes < len(b) {
		return skogul.Error{Source: "net sender", Reason: fmt.Sprintf("Write succeeded, but not all data written. Wrote %d of %d bytes.", nbytes, len(b))}
	}
	return nil
}

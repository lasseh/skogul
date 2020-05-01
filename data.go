/*
 * skogul, core data structures
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

package skogul

import (
	"encoding/json"
	"fmt"
	"time"
)

var dataLog = Logger("core", "data")

/*
Container is the top-level object for one or more metric.

If a Template is provided, it will used as the initial value for each of
the metrics - this is expanded by the transformers.Template transformers,
and internal code does not need to worry about this.

A single Container instance is typically the result of a single POST to the
HTTP receiver or similar.
*/
type Container struct {
	Template *Metric   `json:"template,omitempty"`
	Metrics  []*Metric `json:"metrics"`
}

/*
Metric is a collection of measurements related to the same metadata and
point in time.

A good example of a single metric is port statistics for a single interface
on a router.

Both Metadata and Data is proided. The difference is generally in how data
is accessed. Metadata is data you will search for - e.g. the name of the
router, the name of the interface, etc. It is also possible to add more
dynamic data as metadata, such as OS versions, but exactly how this will be
handled will be up to underlying storage engines. E.g.: for influxdb, anything
in metadata will be an (indexed) tag, so having reasonably rich metadata is
perfectly fine, but you may want to keep an eye on the granularity.

A simple rule of thumb:

Metadata is what you search with.

Data is what you search for.

Example:

	{
		"time": "2019-03-25T12:00:00Z",
		"metadata": {
			"device": "routera",
			"os": "JUNOS 15.4R1",
			"chassisId": "something"
		},
		"data": {
			"uptime": 124125124,
			"cputemp": 22
		}
	}

It is possible to have nested data, however, it is NOT a requirement
that a sender accepts this. And in general, it is better to "flatten"
out data into multiple metrics instead. This can be done with a
(custom) transformer.

Example of a nested structure:

	{
		"time": "2019-03-25T12:00:00Z",
		"metadata": {
			"device": "routera",
			"os": "JUNOS 15.4R1",
			"chassisId": "something"
		},
		"data": {
			"ports": {
				"ge-0/0/0": {
					"ifHCInOctets":  5,
					"ifHCOutOctets": 10
				},
				"ge-0/0/1": {
					"ifHCInOctets":  2,
					"ifHCOutOctets": 20
				}
			}
		}
	}

This is legal, but it's probably wise to use a transformer to change it into:

	{
		"time": "2019-03-25T12:00:00Z",
		"metadata": {
			"device": "routera",
			"os": "JUNOS 15.4R1",
			"chassisId": "something",
			"port": "ge-0/0/0"
		},
		"data": {
			"ifHCInOctets":  5,
			"ifHCOutOctets": 10
			}
		}
	},
	{
		"time": "2019-03-25T12:00:00Z",
		"metadata": {
			"device": "routera",
			"os": "JUNOS 15.4R1",
			"chassisId": "something",
			"port": "ge-0/0/1"
		},
		"data": {
			"ifHCInOctets":  2,
			"ifHCOutOctets": 20
			}
		}
	}
*/
type Metric struct {
	Time     *time.Time             `json:"timestamp,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty"`
}

// Validate a single metric.
func (m *Metric) validate() error {
	if m.Data == nil {
		return Error{Reason: "Missing data for metric"}
	}
	if len(m.Data) <= 0 {
		return Error{Reason: "Missing data for metric"}
	}
	return nil
}

/*
Validate checks the validity of the container, verifying that it follows
the exepcted spec. It should be called by any HTTP receiver after
accepting a Container from an external source. It is NOT required
nor recommended to use Validate in senders - the data is already
validated by that time.
*/
func (c *Container) Validate() error {
	if c.Metrics == nil {
		return Error{Source: "container validation", Reason: "Missing metrics[] data"}
	}
	if len(c.Metrics) <= 0 {
		return Error{Reason: "Empty metrics[] data"}
	}
	for _, m := range c.Metrics {
		if m.Time == nil && (c.Template == nil || c.Template.Time == nil) {
			return Error{Reason: "Missing timestamp in both metric and container"}
		}
		err := m.validate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c Container) String() string {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		dataLog.WithError(err).Error("Unable to marshal JSON")
		return ""
	}
	return fmt.Sprintf("%s", b)
}

/*
Duration provides a wrapper around time.Duration to add JSON marshalling
and unmarshalling. It is used whenever time.Duration needs to be exposed in
configuration.
*/
type Duration struct {
	time.Duration
}

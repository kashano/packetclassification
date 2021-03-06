/*
# GnucleusLAN - LAN-only P2P filesharing
# Pattern attributes: good notsofast notsofast
# Protocol groups: p2p open_source
# Wiki: http://www.protocolinfo.org/wiki/GnucleusLAN
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# This pattern has been tested and is believed to work well.
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "gnucleuslan"
	mypattern.RegularExpresion = "(?i)" + "gnuclear connect/[\\x09-\\x0d -~]*user-agent: gnucleus [\\x09-\\x0d -~]*lan:"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}

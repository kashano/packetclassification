/*
# Medal of Honor Allied Assault - an Electronic Arts game
# Pattern attributes: good veryfast fast
# Protocol groups: game proprietary
# Wiki: http://www.protocolinfo.org/wiki/Medal_of_Honor_Allied_Assault
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# This pattern is written and tested by Krzysztof Maciejewski.
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "mohaa"
	mypattern.RegularExpresion = "(?i)" + "^\\xff\\xff\\xff\\xffgetstatus\\x0a"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}

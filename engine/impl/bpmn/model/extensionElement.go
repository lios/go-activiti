package model

import "encoding/xml"

type ExtensionElement struct {
	ExtensionElementName xml.Name           `xml:"extensionElements"`
	TaskListener         []ActivitiListener `xml:"taskListener"`
}

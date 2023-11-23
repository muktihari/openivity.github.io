package schema

import (
	"encoding/xml"
	"strconv"
)

type Application struct {
	Name       string `xml:"Name"`
	Build      *Build `xml:"Build"`
	LangID     string `xml:"LangID"`     // Specifies the two character ISO 693-1 language id that identifies the installed language of this application. see http://www.loc.gov/standards/iso639-2/ for appropriate ISO identifiers
	PartNumber string `xml:"PartNumber"` // The formatted XXX-XXXXX-XX Garmin part number of a PC application.
}

var _ xml.Unmarshaler = &Application{}

func (a *Application) UnmarshalXML(dec *xml.Decoder, se xml.StartElement) error {
	var targetCharData string
	for {
		token, err := dec.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "Build":
				var build Build
				if err := build.UnmarshalXML(dec, elem); err != nil {
					return err
				}
				a.Build = &build

			default:
				targetCharData = elem.Name.Local
			}
		case xml.CharData:
			switch targetCharData {
			case "Name":
				a.Name = string(elem)
			case "LangID":
				a.LangID = string(elem)
			case "PartNumber":
				a.PartNumber = string(elem)
			}
			targetCharData = ""
		case xml.EndElement:
			if elem == se.End() {
				return nil
			}
		}
	}
}

type Build struct {
	Type    BuildType
	Version *Version
}

var _ xml.Unmarshaler = &Build{}

func (b *Build) UnmarshalXML(dec *xml.Decoder, se xml.StartElement) error {
	var targetCharData string
	for {
		token, err := dec.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "Version":
				var version Version
				if err := version.UnmarshalXML(dec, elem); err != nil {
					return err
				}
				b.Version = &version
			default:
				targetCharData = elem.Name.Local
			}
		case xml.CharData:
			switch targetCharData {
			case "Type":
				b.Type = BuildType(elem)
			}
			targetCharData = ""
		case xml.EndElement:
			if elem == se.End() {
				return nil
			}
		}
	}
}

type BuildType string

const (
	BuildTypeInternal BuildType = "Internal"
	BuildTypeAlpha    BuildType = "Alpha"
	BuildTypeBeta     BuildType = "Beta"
	BuildTypeRelease  BuildType = "Release"
)

type Device struct {
	Name      string   `xml:"Name"`
	UnitId    uint32   `xml:"UnitId"`
	ProductID uint16   `xml:"ProductID"`
	Version   *Version `xml:"Version"`
}

var _ xml.Unmarshaler = &Device{}

func (d *Device) UnmarshalXML(dec *xml.Decoder, se xml.StartElement) error {
	var targetCharData string
	for {
		token, err := dec.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "Version":
				var version Version
				if err := version.UnmarshalXML(dec, elem); err != nil {
					return err
				}
				d.Version = &version

			default:
				targetCharData = elem.Name.Local
			}
		case xml.CharData:
			switch targetCharData {
			case "Name":
				d.Name = string(elem)
			case "UnitId":
				u, err := strconv.ParseUint(string(elem), 10, 32)
				if err != nil {
					return err
				}
				d.UnitId = uint32(u)
			case "ProductID":
				u, err := strconv.ParseUint(string(elem), 10, 16)
				if err != nil {
					return err
				}
				d.ProductID = uint16(u)
			}
			targetCharData = ""
		case xml.EndElement:
			if elem == se.End() {
				return nil
			}
		}
	}
}

type Version struct {
	VersionMajor uint16 `xml:"VersionMajor"`
	VersionMinor uint16 `xml:"VersionMinor"`
	BuildMajor   uint16 `xml:"BuildMajor"`
	BuildMinor   uint16 `xml:"BuildMinor"`
}

var _ xml.Unmarshaler = &Version{}

func (v *Version) UnmarshalXML(dec *xml.Decoder, se xml.StartElement) error {
	var targetCharData string
	for {
		token, err := dec.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			targetCharData = elem.Name.Local
		case xml.CharData:
			switch targetCharData {
			case "VersionMajor":
				u, err := strconv.ParseUint(string(elem), 10, 16)
				if err != nil {
					return err
				}
				v.VersionMajor = uint16(u)
			case "VersionMinor":
				u, err := strconv.ParseUint(string(elem), 10, 16)
				if err != nil {
					return err
				}
				v.VersionMinor = uint16(u)
			case "BuildMajor":
				u, err := strconv.ParseUint(string(elem), 10, 16)
				if err != nil {
					return err
				}
				v.BuildMajor = uint16(u)
			case "BuildMinor":
				u, err := strconv.ParseUint(string(elem), 10, 16)
				if err != nil {
					return err
				}
				v.BuildMinor = uint16(u)
			}
			targetCharData = ""
		case xml.EndElement:
			if elem == se.End() {
				return nil
			}
		}
	}
}

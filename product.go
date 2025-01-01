package ndcd

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Product struct {
	ProprietaryName       string
	ProprietaryNameSuffix string
	NonProprietaryName    string
	LabelerName           string
	Labeler               *Labeler
	SubstanceName         string
	Variations            []Variation
}

type Variation struct {
	ProductID             string
	ProductNDC            string
	DosageFormName        []string
	RouteName             string
	ProductTypeName       string
	MarketingCategoryName string
	ApplicationNumber     string
	StartMarketingDate    time.Time
	EndMarketingDate      time.Time
	PharmClasses          []string
	DEASchedule           string
	StrengthNumber        string
	StrengthUnit          string
}

func parseLine(rawStr string) Product {
	var p Product
	rawStringArray := strings.Split(rawStr, "	")

	p.ProprietaryName = rawStringArray[3]
	p.ProprietaryNameSuffix = rawStringArray[4]
	p.NonProprietaryName = rawStringArray[5]
	p.LabelerName = func() string {
		str := strings.ToLower(rawStringArray[12])
		str = strings.Title(str)
		str += " "
		str = strings.Replace(str, ",", "", -1)
		str = strings.Replace(str, ".", "", -1)
		str = strings.Replace(str, "Inc ", "INC. ", -1)
		str = strings.Replace(str, "Llc ", "LLC. ", -1)
		str = strings.Replace(str, "Ltd ", "LTD. ", -1)
		str = strings.Replace(str, " Of ", " of ", -1)
		str = strings.Replace(str, "Usa ", "USA ", -1)
		str = strings.Replace(str, "Pty ", "PTY. ", -1)
		str = strings.Replace(str, "Llc ", "LLC. ", -1)
		str = strings.Replace(str, "Pvt ", "PVT. ", -1)
		str = strings.Replace(str, "PTY ", "PTY. ", -1)
		if !strings.Contains(str, "Corporation ") {
			str = strings.Replace(str, "Corp ", "CORP. ", -1)
		}
		str = strings.TrimSpace(str)
		return str
	}()
	p.SubstanceName = rawStringArray[13]

	var v Variation

	v.DosageFormName = strings.Split(rawStringArray[6], ",")
	v.RouteName = rawStringArray[7]
	v.ProductID = rawStringArray[0]
	v.ProductNDC = rawStringArray[1]
	v.ProductTypeName = rawStringArray[2]
	if len(rawStringArray[8]) == 8 {
		smdRaw := rawStringArray[8]
		smdY, _ := strconv.ParseInt(smdRaw[:4], 10, 64)
		smdM, _ := strconv.ParseInt(smdRaw[4:6], 10, 64)
		smdD, _ := strconv.ParseInt(smdRaw[6:], 10, 64)
		smd := time.Date(int(smdY), time.Month(smdM), int(smdD), 0, 0, 0, 0, &time.Location{})
		v.StartMarketingDate = smd
	}
	if len(rawStringArray[9]) == 8 {
		emdRaw := rawStringArray[9]
		emdY, _ := strconv.ParseInt(emdRaw[:4], 10, 64)
		emdM, _ := strconv.ParseInt(emdRaw[4:6], 10, 64)
		emdD, _ := strconv.ParseInt(emdRaw[6:], 10, 64)
		emd := time.Date(int(emdY), time.Month(emdM), int(emdD), 0, 0, 0, 0, &time.Location{})
		v.EndMarketingDate = emd
	} else {
		v.EndMarketingDate = time.Time{}
	}
	v.MarketingCategoryName = rawStringArray[10]
	v.ApplicationNumber = rawStringArray[11]
	v.StrengthNumber = rawStringArray[14]
	v.StrengthUnit = rawStringArray[15]
	if rawStringArray[16] != "" {
		v.PharmClasses = strings.Split(rawStringArray[16], ",")
	}
	if rawStringArray[17] != "" {
		v.DEASchedule = strings.TrimRight(rawStringArray[17], "\r\n")
	}
	p.Variations = []Variation{v}

	return p
}

func (p *Product) String() string {
	return fmt.Sprintf("{ %s | %d variations }", p.ProprietaryName, len(p.Variations))
}

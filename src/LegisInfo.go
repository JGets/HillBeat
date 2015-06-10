package main

import (
    "fmt"
    "encoding/xml"
    "net/http"
)

const (
    LangEn = 0
    LangFr = 1
)

type BillTitle struct {
    LanguageStr string `xml:"language,attr"`
    Language int
    Value string `xml:",chardata"`
}

type BillBiTitle struct {
    Title []BillTitle `xml:"Title"`
}

type BillPerson struct {
    IdStr string `xml:"id,attr"`
    Id uint
    Gender string `xml:"Gender,attr"`
    FullName string `xml:"FullName"`
    FirstName string `xml:"FirstName"`
    MiddleName string `xml:"MiddleName"`
    LastName string `xml:"LastName"`
}


type BillPoliticalParty struct {
    Title []BillTitle `xml:"Title"`
    Abbreviation BillBiTitle `xml:"abbreviation"`
}

type BillAffiliation struct {
    IdStr string `xml:"id,attr"`
    Id uint
    Title []BillTitle `xml:"Title"`
    Person BillPerson `xml:"Person"`
    Party BillPoliticalParty `xml:"PoliticalParty"`
}

type BillStatute struct {
    Year int `xml:"Year"`
    Chapter int `xml:"Chapter"`
} 

type BillPublicationFile struct {
    LanguageStr string `xml:"language,attr"`
    Language int
    RelativePath string `xml:"relativePath,attr"`
}

type BillPublicationFiles struct {
    Files []BillPublicationFile `xml:"PublicationFile"`
}

type BillPublication struct {
    IdStr string `xml:"id,attr"`
    Id uint
    Title []BillTitle `xml:"Title"`
    PublicationFiles BillPublicationFiles `xml:"PublicationFiles"`
}

type BillPublications struct {
    Publications []BillPublication `xml:"Publication"`
}

type BillEvent struct {
    IdStr string `xml:"id,attr"`
    Id uint
    Chamber string `xml:"chamber,attr"`
    DateStr string `xml:"date,attr"`
    MeetingNumberStr string `xml:"meetingNumber,attr"`
    MeetingNumber uint
    Status BillBiTitle `xml:"Status"`
    //Description
}

type BillLastEvent struct {
    Event BillEvent `xml:"Event"`
    Progress string `xml:"Progress"`
}

type BillLegislativeEvents struct {
    Events []BillEvent `xml:"Event"`
}

type BillEvents struct {
    CurrentStage string `xml:"laagCurrentStage,attr"`
    LastMajorStageEvent BillLastEvent `xml:"LastMajorStageEvent"`
    Events BillLegislativeEvents `xml:"LegislativeEvents"`
}

type BillParliamentSession struct {
    ParliamentNumberStr string `xml:"parliamentNumber,attr"`
    ParliamentNumber uint
    SessionNumberStr string `xml:"sessionNumber,attr"`
    SessionNumber uint
}

type BillNumber struct {
    Prefix string `xml:"prefix,attr"`
    Number string `xml:"number,attr"`
}

type Bill struct {
    IdStr string `xml:"id,attr"`
    Id uint
    LastUpdatedStr string `xml:"lastUpdated,attr"`
    IntroducedDateStr string `xml:"BillIntroducedDate"`
    Session BillParliamentSession `xml:"ParliamentSession"`
    Number BillNumber `xml:"BillNumber"`
    LongTitle BillBiTitle `xml:"BillTitle"`
    ShortTitle BillBiTitle `xml:"ShortTitle"`
    Type BillBiTitle `xml:"BillType"`
    SponsorAffiliation BillAffiliation `xml:"SponsorAffiliation"`
    PrimeMinister BillAffiliation `xml:"PrimeMinister"`
    Statute BillStatute `xml:"Statute"`
    Publications BillPublications `xml:"Publications"`
    Events BillEvents `xml:"Events"`
}

type Bills struct {
    BillList []Bill `xml:"Bill"`
}

func main() {
    var bills Bills
    resp, err := http.Get("http://www.parl.gc.ca/LegisInfo/Home.aspx?Language=E&Mode=1&ParliamentSession=41-2&download=xml")
    if err != nil {
        //ERROR
    }

    decoder := xml.NewDecoder(resp.Body)
    for {
        t, _ := decoder.Token()
        if t == nil {
            break
        }

        switch se := t.(type) {
        case xml.StartElement:
            if se.Name.Local == "Bills" {
                decoder.DecodeElement(&bills, &se)
            }
        default:
        }
    }

    fmt.Printf("Total number of bills: %d \n", len(bills.BillList))
}

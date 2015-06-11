package LegisInfo

import (
    "fmt"
    "encoding/xml"
    "net/http"
)

type XmlBillTitle struct {
    Language string `xml:"language,attr"`
    Value string `xml:",chardata"`
}

type XmlBillBiTitle struct {
    Title []XmlBillTitle `xml:"Title"`
}

type XmlBillPerson struct {
    Id string `xml:"id,attr"`
    Gender string `xml:"Gender,attr"`
    FullName string `xml:"FullName"`
    FirstName string `xml:"FirstName"`
    MiddleName string `xml:"MiddleName"`
    LastName string `xml:"LastName"`
}


type XmlBillPoliticalParty struct {
    Title []XmlBillTitle `xml:"Title"`
    Abbreviation XmlBillBiTitle `xml:"abbreviation"`
}

type XmlBillAffiliation struct {
    Id string `xml:"id,attr"`
    Title []XmlBillTitle `xml:"Title"`
    Person XmlBillPerson `xml:"Person"`
    Party XmlBillPoliticalParty `xml:"PoliticalParty"`
}

type XmlBillStatute struct {
    Year int `xml:"Year"`
    Chapter int `xml:"Chapter"`
} 

type XmlBillPublicationFile struct {
    Language string `xml:"language,attr"`
    RelativePath string `xml:"relativePath,attr"`
}

type XmlBillPublicationFiles struct {
    Files []XmlBillPublicationFile `xml:"PublicationFile"`
}

type XmlBillPublication struct {
    Id string `xml:"id,attr"`
    Title []XmlBillTitle `xml:"Title"`
    PublicationFiles XmlBillPublicationFiles `xml:"PublicationFiles"`
}

type XmlBillPublications struct {
    Publications []XmlBillPublication `xml:"Publication"`
}

type XmlBillEvent struct {
    Id string `xml:"id,attr"`
    Chamber string `xml:"chamber,attr"`
    DateStr string `xml:"date,attr"`
    MeetingNumberStr string `xml:"meetingNumber,attr"`
    MeetingNumber uint
    Status XmlBillBiTitle `xml:"Status"`
    //Description
}

type XmlBillLastEvent struct {
    Event XmlBillEvent `xml:"Event"`
    Progress string `xml:"Progress"`
}

type XmlBillLegislativeEvents struct {
    Events []XmlBillEvent `xml:"Event"`
}

type XmlBillEvents struct {
    CurrentStage string `xml:"laagCurrentStage,attr"`
    LastMajorStageEvent XmlBillLastEvent `xml:"LastMajorStageEvent"`
    Events XmlBillLegislativeEvents `xml:"LegislativeEvents"`
}

type XmlBillParliamentSession struct {
    ParliamentNumber string `xml:"parliamentNumber,attr"`
    SessionNumber string `xml:"sessionNumber,attr"`
}

type XmlBillNumber struct {
    Prefix string `xml:"prefix,attr"`
    Number string `xml:"number,attr"`
}

type XmlBill struct {
    Id string `xml:"id,attr"`
    LastUpdated string `xml:"lastUpdated,attr"`
    IntroducedDate string `xml:"BillIntroducedDate"`
    Session XmlBillParliamentSession `xml:"ParliamentSession"`
    Number XmlBillNumber `xml:"BillNumber"`
    LongTitle XmlBillBiTitle `xml:"BillTitle"`
    ShortTitle XmlBillBiTitle `xml:"ShortTitle"`
    Type XmlBillBiTitle `xml:"BillType"`
    SponsorAffiliation XmlBillAffiliation `xml:"SponsorAffiliation"`
    PrimeMinister XmlBillAffiliation `xml:"PrimeMinister"`
    Statute XmlBillStatute `xml:"Statute"`
    Publications XmlBillPublications `xml:"Publications"`
    Events XmlBillEvents `xml:"Events"`
}

type XmlBills struct {
    BillList []XmlBill `xml:"Bill"`
}

func GetLegisInfoXmlAllBills() XmlBills {
    var bills XmlBills
    var page = 0
    var prevLength = -1

    //Increment page # until parsed length doesn't change; 
    //  limit to 10 pages (to prevent infinite loop; and will we really ever have more than 5000 bills in a session?)
    for prevLength < len(bills.BillList) && page < 10 {
        prevLength = len(bills.BillList)
        page++

        resp, err := http.Get(fmt.Sprintf("http://www.parl.gc.ca/LegisInfo/Home.aspx?Language=E&Mode=1&ParliamentSession=41-2&download=xml&page=%d", page))
        if err != nil {
            panic(err)
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

        //fmt.Printf("read number of bills: %d, page %d\n", len(bills.BillList), page)
    }

    //fmt.Printf("Total number of bills: %d \tTested %d pages\n\n", len(bills.BillList), page)

    return bills
}

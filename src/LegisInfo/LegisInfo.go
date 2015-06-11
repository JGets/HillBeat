package LegisInfo

import (
    "fmt"
    "./xml"
    "time"
)

type Person struct {
    Id int

    //Do we really need full and part?
    FullName string
    FirstName string
    MiddleName string
    LastNamee string
}

type Party struct {
    TitleEnglish string
    TitleFrench string
    AbbrEnglish string
    AbbrFrench string
}

type BillSponsor struct {
    Id int

    TitleEnglish string
    TitleFrench string

    Person Person
    Party Party
}

type PublicationFile struct {
    Language string
    Path string
}

type Publication struct {
    Id int

    TitleEnglish string
    TitleFrench string
    Files []PublicationFile
}

type Event struct {
    Id int

    Chamber string
    Date Time
    MeetingNumber int
    //Committe Meetings???

    TitleEnglish string
    TitleFrench string

}

type Bill struct {
    Id int
    
    LastUpdated Time
    Introduced Time
    
    //Parliament info
    Parliament int
    Session int
    
    //Number info
    NumberPrefix byte //char
    Number int

    //Title
    LongTitleEnglish string
    LongTitleFrench string
    ShortTitleEnglish string
    ShortTitleFrench string

    //Type
    TypeEnglish string
    TypeFrench string

    Sponsor BillSponsor
    PrimeMinister BillSponsor

    //Statute
    StatuteYear int
    StatuteChapter int

    //Publications
    Publications []Publication

    //Events
    Stage string
    Progress string
    Events []Event
}

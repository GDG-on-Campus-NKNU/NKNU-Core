package sso

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type course struct {
	Name        string `json:"name"`
	Credit      string `json:"credit"`
	Category    string `json:"category"`
	Requirement string `json:"requirement"`
	Score       string `json:"score"`
}

type historyScore struct {
	Title           string     `json:"title"`
	Courses         *[]*course `json:"courses"`
	AverageScore    string     `json:"averageScore"`
	ConductScore    string     `json:"conductScore"`
	Credits         string     `json:"credits"`
	EarnedCredits   string     `json:"earnedCredits"`
	SemesterRanking string     `json:"semesterRanking"`
	ClassSize       string     `json:"classSize"`
}

func getSplitter(s string) string {
	if strings.Contains(s, "/") {
		return " / "
	} else {
		return " ／ "
	}
}

func GetHistoryScore(session *Session) (*[]*historyScore, error) {
	bodyString, err := newRequest("GET", "https://sso.nknu.edu.tw/StudentProfile/Day/CourseScoreAll.aspx", nil, session.AspNETSessionId, nil)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(bodyString))
	if err != nil {
		return nil, err
	}

	var historyScoreList []*historyScore
	doc.Find("div[id=ctl00_phMain_divRow]").Find(`div[style="width:100%;"]`).Each(func(i int, s *goquery.Selection) {
		title := s.Find("span").Text()
		blocks := s.Find("div[class=row]")
		// courses
		var courseList []*course
		blocks.Eq(0).Find("tr").Each(func(i int, s *goquery.Selection) {
			tds := s.Find("td")
			if tds.Length() == 0 {
				return
			}

			courseName := tds.Eq(0).Find("a").Text()
			courseCredit := tds.Eq(1).Text()
			courseCategory := tds.Eq(2).Text()
			courseRequirement := tds.Eq(3).Text()
			courseScore := tds.Eq(4).Text()
			courseList = append(courseList, &course{
				Name:        courseName,
				Credit:      courseCredit,
				Category:    courseCategory,
				Requirement: courseRequirement,
				Score:       courseScore,
			})
		})

		var avgScore string
		var conductScore string
		var credits string
		var earnedCredits string
		var semesterRanking string
		var classSize string

		if blocks.Eq(1).Find("table").Children().Length() > 0 {
			scoreF := blocks.Eq(1).Find("tr")
			scoreRowContent := scoreF.Eq(0).Find(`td[colspan="6"]`).Text()
			scoreRow := strings.Split(scoreRowContent, getSplitter(scoreRowContent))
			avgScore, conductScore = scoreRow[0], scoreRow[1]
			creditRowContent := scoreF.Eq(1).Find(`td[colspan="6"]`).Text()
			creditRow := strings.Split(creditRowContent, getSplitter(creditRowContent))
			credits, earnedCredits = creditRow[0], creditRow[1]
			rankingRowContent := scoreF.Eq(2).Find(`td[colspan="6"]`).Text()
			rankingRow := strings.Split(rankingRowContent, getSplitter(rankingRowContent))
			semesterRanking, classSize = rankingRow[0], rankingRow[1]
		}

		historyScoreList = append(historyScoreList, &historyScore{
			Title:           title,
			Courses:         &courseList,
			AverageScore:    avgScore,
			ConductScore:    conductScore,
			Credits:         credits,
			EarnedCredits:   earnedCredits,
			SemesterRanking: semesterRanking,
			ClassSize:       classSize,
		})
	})

	return &historyScoreList, nil
}

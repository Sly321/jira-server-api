package issue

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jira-server-api/main/pkg/client/constants"
	"jira-server-api/main/pkg/util/logging"
	"jira-server-api/main/pkg/util/rest"
	"log"
)

type Issue struct {
	Expand string `json:"expand"`
	ID     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields struct {
		Customfield21323 []struct {
			IssueKey       string  `json:"issueKey"`
			Status         string  `json:"status"`
			StatusStyle    string  `json:"statusStyle"`
			Ok             int     `json:"ok"`
			OkPercent      float64 `json:"okPercent"`
			OkJql          string  `json:"okJql"`
			Nok            int     `json:"nok"`
			NokPercent     float64 `json:"nokPercent"`
			NokJql         string  `json:"nokJql"`
			Notrun         int     `json:"notrun"`
			NotrunPercent  float64 `json:"notrunPercent"`
			NotRunJql      string  `json:"notRunJql"`
			Unknown        int     `json:"unknown"`
			UnknownPercent float64 `json:"unknownPercent"`
			UnknownJql     string  `json:"unknownJql"`
		} `json:"customfield_21323"`
		Issuetype struct {
			Self        string `json:"self"`
			ID          string `json:"id"`
			Description string `json:"description"`
			IconURL     string `json:"iconUrl"`
			Name        string `json:"name"`
			Subtask     bool   `json:"subtask"`
			AvatarID    int    `json:"avatarId"`
		} `json:"issuetype"`
		Customfield19004 interface{} `json:"customfield_19004"`
		Customfield23900 interface{} `json:"customfield_23900"`
		Customfield23222 interface{} `json:"customfield_23222"`
		Customfield17100 interface{} `json:"customfield_17100"`
		Customfield23220 interface{} `json:"customfield_23220"`
		Timespent        int         `json:"timespent"`
		Customfield13100 string      `json:"customfield_13100"`
		Project          struct {
			Self           string `json:"self"`
			ID             string `json:"id"`
			Key            string `json:"key"`
			Name           string `json:"name"`
			ProjectTypeKey string `json:"projectTypeKey"`
			AvatarUrls     struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
		} `json:"project"`
		FixVersions        []interface{} `json:"fixVersions"`
		Aggregatetimespent int           `json:"aggregatetimespent"`
		Resolution         interface{}   `json:"resolution"`
		Customfield19800   interface{}   `json:"customfield_19800"`
		Resolutiondate     interface{}   `json:"resolutiondate"`
		Customfield22804   interface{}   `json:"customfield_22804"`
		Workratio          int           `json:"workratio"`
		Customfield23219   interface{}   `json:"customfield_23219"`
		Customfield22600   interface{}   `json:"customfield_22600"`
		Customfield24800   interface{}   `json:"customfield_24800"`
		Customfield24801   interface{}   `json:"customfield_24801"`
		LastViewed         string        `json:"lastViewed"`
		Watches            struct {
			Self       string `json:"self"`
			WatchCount int    `json:"watchCount"`
			IsWatching bool   `json:"isWatching"`
		} `json:"watches"`
		Customfield24802 interface{} `json:"customfield_24802"`
		Customfield22200 interface{} `json:"customfield_22200"`
		Customfield24803 interface{} `json:"customfield_24803"`
		Customfield24804 interface{} `json:"customfield_24804"`
		Customfield24805 interface{} `json:"customfield_24805"`
		Created          string      `json:"created"`
		Priority         struct {
			Self    string `json:"self"`
			IconURL string `json:"iconUrl"`
			Name    string `json:"name"`
			ID      string `json:"id"`
		} `json:"priority"`
		Customfield10100              interface{}   `json:"customfield_10100"`
		Customfield16800              string        `json:"customfield_16800"`
		Customfield14500              interface{}   `json:"customfield_14500"`
		Labels                        []interface{} `json:"labels"`
		Aggregatetimeoriginalestimate int           `json:"aggregatetimeoriginalestimate"`
		Timeestimate                  int           `json:"timeestimate"`
		Versions                      []interface{} `json:"versions"`
		Issuelinks                    []interface{} `json:"issuelinks"`
		Customfield25300              string        `json:"customfield_25300"`
		Customfield25301              string        `json:"customfield_25301"`
		Assignee                      struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
			TimeZone    string `json:"timeZone"`
		} `json:"assignee"`
		Updated string `json:"updated"`
		Status  struct {
			Self           string `json:"self"`
			Description    string `json:"description"`
			IconURL        string `json:"iconUrl"`
			Name           string `json:"name"`
			ID             string `json:"id"`
			StatusCategory struct {
				Self      string `json:"self"`
				ID        int    `json:"id"`
				Key       string `json:"key"`
				ColorName string `json:"colorName"`
				Name      string `json:"name"`
			} `json:"statusCategory"`
		} `json:"status"`
		Components           []interface{} `json:"components"`
		Customfield19500     interface{}   `json:"customfield_19500"`
		Customfield23201     interface{}   `json:"customfield_23201"`
		Customfield17201     interface{}   `json:"customfield_17201"`
		Customfield23200     interface{}   `json:"customfield_23200"`
		Customfield17200     interface{}   `json:"customfield_17200"`
		Timeoriginalestimate int           `json:"timeoriginalestimate"`
		Description          string        `json:"description"`
		// Customfield13002     []time.Time   `json:"customfield_13002"`
		Customfield15503 interface{} `json:"customfield_15503"`
		Customfield15622 interface{} `json:"customfield_15622"`
		Customfield15501 interface{} `json:"customfield_15501"`
		Timetracking     struct {
			OriginalEstimate         string `json:"originalEstimate"`
			RemainingEstimate        string `json:"remainingEstimate"`
			TimeSpent                string `json:"timeSpent"`
			OriginalEstimateSeconds  int    `json:"originalEstimateSeconds"`
			RemainingEstimateSeconds int    `json:"remainingEstimateSeconds"`
			TimeSpentSeconds         int    `json:"timeSpentSeconds"`
		} `json:"timetracking"`
		Customfield15623      interface{}   `json:"customfield_15623"`
		Customfield15502      interface{}   `json:"customfield_15502"`
		Customfield15700      interface{}   `json:"customfield_15700"`
		Attachment            []interface{} `json:"attachment"`
		Aggregatetimeestimate int           `json:"aggregatetimeestimate"`
		Customfield20606      interface{}   `json:"customfield_20606"`
		Customfield20604      interface{}   `json:"customfield_20604"`
		Customfield20605      interface{}   `json:"customfield_20605"`
		Customfield22507      interface{}   `json:"customfield_22507"`
		Customfield20800      []interface{} `json:"customfield_20800"`
		Customfield20602      interface{}   `json:"customfield_20602"`
		Customfield22506      interface{}   `json:"customfield_22506"`
		Customfield20801      interface{}   `json:"customfield_20801"`
		Customfield20603      interface{}   `json:"customfield_20603"`
		Customfield22503      interface{}   `json:"customfield_22503"`
		Customfield24500      interface{}   `json:"customfield_24500"`
		// Customfield20400      time.Time     `json:"customfield_20400"`
		Customfield24501 interface{} `json:"customfield_24501"`
		Customfield20005 struct {
			Self     string `json:"self"`
			Value    string `json:"value"`
			ID       string `json:"id"`
			Disabled bool   `json:"disabled"`
		} `json:"customfield_20005"`
		Customfield20401 interface{}   `json:"customfield_20401"`
		Summary          string        `json:"summary"`
		Customfield18200 []interface{} `json:"customfield_18200"`
		Customfield24900 interface{}   `json:"customfield_24900"`
		Customfield18201 interface{}   `json:"customfield_18201"`
		Creator          struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
			TimeZone    string `json:"timeZone"`
		} `json:"creator"`
		Customfield18202 interface{}   `json:"customfield_18202"`
		Customfield16100 interface{}   `json:"customfield_16100"`
		Subtasks         []interface{} `json:"subtasks"`
		Reporter         struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
			TimeZone    string `json:"timeZone"`
		} `json:"reporter"`
		Aggregateprogress struct {
			Progress int `json:"progress"`
			Total    int `json:"total"`
			Percent  int `json:"percent"`
		} `json:"aggregateprogress"`
		Customfield10004 interface{} `json:"customfield_10004"`
		Customfield16700 interface{} `json:"customfield_16700"`
		Environment      interface{} `json:"environment"`
		Duedate          interface{} `json:"duedate"`
		Customfield25000 interface{} `json:"customfield_25000"`
		Progress         struct {
			Progress int `json:"progress"`
			Total    int `json:"total"`
			Percent  int `json:"percent"`
		} `json:"progress"`
		Comment struct {
			Comments []struct {
				Self   string `json:"self"`
				ID     string `json:"id"`
				Author struct {
					Self         string `json:"self"`
					Name         string `json:"name"`
					Key          string `json:"key"`
					EmailAddress string `json:"emailAddress"`
					AvatarUrls   struct {
						Four8X48  string `json:"48x48"`
						Two4X24   string `json:"24x24"`
						One6X16   string `json:"16x16"`
						Three2X32 string `json:"32x32"`
					} `json:"avatarUrls"`
					DisplayName string `json:"displayName"`
					Active      bool   `json:"active"`
					TimeZone    string `json:"timeZone"`
				} `json:"author"`
				Body         string `json:"body"`
				UpdateAuthor struct {
					Self         string `json:"self"`
					Name         string `json:"name"`
					Key          string `json:"key"`
					EmailAddress string `json:"emailAddress"`
					AvatarUrls   struct {
						Four8X48  string `json:"48x48"`
						Two4X24   string `json:"24x24"`
						One6X16   string `json:"16x16"`
						Three2X32 string `json:"32x32"`
					} `json:"avatarUrls"`
					DisplayName string `json:"displayName"`
					Active      bool   `json:"active"`
					TimeZone    string `json:"timeZone"`
				} `json:"updateAuthor"`
				Created string `json:"created"`
				Updated string `json:"updated"`
			} `json:"comments"`
			MaxResults int `json:"maxResults"`
			Total      int `json:"total"`
			StartAt    int `json:"startAt"`
		} `json:"comment"`
		Worklog struct {
			StartAt    int `json:"startAt"`
			MaxResults int `json:"maxResults"`
			Total      int `json:"total"`
			Worklogs   []struct {
				Self   string `json:"self"`
				Author struct {
					Self         string `json:"self"`
					Name         string `json:"name"`
					Key          string `json:"key"`
					EmailAddress string `json:"emailAddress"`
					AvatarUrls   struct {
						Four8X48  string `json:"48x48"`
						Two4X24   string `json:"24x24"`
						One6X16   string `json:"16x16"`
						Three2X32 string `json:"32x32"`
					} `json:"avatarUrls"`
					DisplayName string `json:"displayName"`
					Active      bool   `json:"active"`
					TimeZone    string `json:"timeZone"`
				} `json:"author"`
				UpdateAuthor struct {
					Self         string `json:"self"`
					Name         string `json:"name"`
					Key          string `json:"key"`
					EmailAddress string `json:"emailAddress"`
					AvatarUrls   struct {
						Four8X48  string `json:"48x48"`
						Two4X24   string `json:"24x24"`
						One6X16   string `json:"16x16"`
						Three2X32 string `json:"32x32"`
					} `json:"avatarUrls"`
					DisplayName string `json:"displayName"`
					Active      bool   `json:"active"`
					TimeZone    string `json:"timeZone"`
				} `json:"updateAuthor"`
				Comment          string `json:"comment"`
				Created          string `json:"created"`
				Updated          string `json:"updated"`
				Started          string `json:"started"`
				TimeSpent        string `json:"timeSpent"`
				TimeSpentSeconds int    `json:"timeSpentSeconds"`
				ID               string `json:"id"`
				IssueID          string `json:"issueId"`
			} `json:"worklogs"`
		} `json:"worklog"`
	} `json:"fields"`
}

const ISSUE_API_PATH = "/issue"

func Get(key string) *Issue {
	url := constants.JIRA_REST_URL + ISSUE_API_PATH + "/" + key
	response, err := rest.Get(url)
	logging.D("issue.Get - response.Status: " + response.Status)

	if err != nil {
		log.Fatal(err.Error())
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	response.Body.Close()

	var result Issue

	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("Can not unmarshal json")
		log.Fatal(err)
	}

	return &result
}

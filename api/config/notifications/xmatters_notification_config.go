package notifications

import "github.com/dtcookie/gojson"

// XMattersNotificationConfig Configuration of the xMatters notification.
type XMattersNotificationConfig struct {
	BaseNotificationConfig `json:",type=XMATTERS"`
	URL                    string       `json:"url"`                  // The URL of the xMatters WebHook.
	AcceptAnyCertificate   bool         `json:"acceptAnyCertificate"` // Accept any, including self-signed and invalid, SSL certificate (`true`) or only trusted (`false`) certificates.
	Headers                []HTTPHeader `json:"headers,omitempty"`    // A list of the additional HTTP headers.
	Payload                string       `json:"payload"`              // The content of the message.  You can use the following placeholders:  * `{ImpactedEntities}`: Details about the entities impacted by the problem in form of a JSON array.  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsHTML}`: All problem event details, including root cause, as an HTML-formatted string.  * `{ProblemDetailsJSON}`: All problem event details, including root cause, as a JSON object.  * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
}

// UnmarshalJSON provides custom JSON deserialization
func (xmnc *XMattersNotificationConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, xmnc)
}

// MarshalJSON provides custom JSON serialization
func (xmnc *XMattersNotificationConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(xmnc)
}

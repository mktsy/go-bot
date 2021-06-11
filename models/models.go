package models

type InputMessage struct {
	Object string `json:"id"`
	Entry  []struct {
		ID        string `json:"id"`
		Time      int64  `json:"time"`
		Messaging []struct {
			Postback struct {
				Title   string `json:"title"`
				Payload string `json:"payload"`
			} `json:"postback"`
			Sender struct {
				ID string `json:"id"`
			} `json:"sender"`
			Recipient struct {
				ID string `json:"id"`
			} `json:"recipient"`
			Timestamp int64 `json:"timestamp"`
			Message   struct {
				Mid  string `json:"mid"`
				Text string `json:"text"`
				Nlp  struct {
					Entities struct {
						Sentiment []struct {
							Confidence float64 `json:"confidence"`
							Value      string  `json:"value"`
						} `json:"sentiment"`
						Greeting []struct {
							Confidence float64 `json:"confidence"`
							Value      string  `json:"value"`
						} `json:"greeting"`
						Email []struct {
							Confidence float64 `json:"confidence"`
							Value      string  `json:"value"`
						} `json:"email"`
					} `json:"entities"`
					DetectedLocales []struct {
						Locale     string  `json:"locale"`
						Confidence float64 `json:"confidence"`
					} `json:"detected_localse"`
				} `json:"nlp"`
				QuickReply struct {
					Playload string `json:"payload"`
				} `json:"quick_play"`
			} `json:"message"`
		} `json:"messaging"`
	} `json:"entry"`
}

type Recipient struct {
	ID string `json:"id"`
}

type QuickReply struct {
	ContentType string `json:"content_type,omitempty"`
	Title       string `json:"title,omitempty"`
	Payload     string `json:"payload,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
}

type Message struct {
	Text         string       `json:"text,omitempty"`
	Mid          string       `json:"mid,omitempty"`
	QuickReplies []QuickReply `json:"quick_replies"`
}

type Button struct {
	Type    string `json:"type,omitempty"`
	Title   string `json:"title,omitempty"`
	Payload string `json:"payload,omitempty"`
	URL     string `json:"url,omitempty"`
}

type Element struct {
	Title         string        `json:"title,omitempty"`
	Subtitle      string        `json:"subtitle,omitempty"`
	ImageURL      string        `json:"image_url,omitempty"`
	DefaultAction DefaultAction `json:"default_action,omitempty"`
	Buttons       []Button      `json:"buttons,omitempty"`
}

type DefaultAction struct {
	Type                string `json:"type,omitempty"`
	URL                 string `json:"url,omitempty"`
	WebViewHeightRation string `json:"webview_height_ratio,omitempty"`
}

type Attachment struct {
	Attachment struct {
		Type    string `json:"type,omitempty"`
		Payload struct {
			TemplateType string    `json:"template_type,omitempty"`
			Elements     []Element `json:"elements,omitempty"`
		} `json:"payload,omitempty"`
	} `json:"attachment,omitempty"`
}

type ResponseAttachment struct {
	Recipient   Recipient  `json:"recipient"`
	MessageType string     `json:"message_type,omitempty"`
	Message     Attachment `json:"message,omitempty"`
}

type ResponseMessage struct {
	Recipient     Recipient `json:"recipient"`
	MessagingType string    `json:"messaging_type,omitempty"`
	Message       Message   `json:"message,omitempty"`
}

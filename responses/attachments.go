package responses

// SearchAttachmentSubscribers is common attachment with subscribers information
// for *.search API methods.
type SearchAttachmentSubscribers struct {
	SubscriberPHIDs    []string `json:"subscriberPHIDs"`
	SubscriberCount    int      `json:"subscriberCount"`
	ViewerIsSubscribed bool     `json:"viewerIsSubscribed"`
}

// SearchAttachmentProjects is common attachment with projects information for
// *.search API methods.
type SearchAttachmentProjects struct {
	ProjectPHIDs []string `json:"projectPHIDs"`
}

// SearchAttachmentReviewers is attachment with revision reviewers information
// for differenial.revision.search API method.
type SearchAttachmentReviewers struct {
	Reviewers []AttachmentReviewer `json:"reviewers"`
}

// AttachmentReviewer is a single revision reviewer in reviewers list.
type AttachmentReviewer struct {
	ReviewerPHID    string `json:"reviewerPHID"`
	Status          string `json:"status"`
	IsBlocking      bool   `json:"isBlocking"`
	ActorPHID       string `json:"actorPHID"`
	IsCurrentAction bool   `json:"isCurrentAction"`
}

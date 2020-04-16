package e621

type File struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Ext    string `json:"ext"`
	Size   int    `json:"size"`
	Md5    string `json:"md5"`
	URL    string `json:"url"`
	Has    bool   `json:"has"`
}

type Score struct {
	Up    int `json:"up"`
	Down  int `json:"down"`
	Total int `json:"total"`
}

type TagContainer struct {
	General   []string `json:"general"`
	Species   []string `json:"species"`
	Character []string `json:"character"`
	Copyright []string `json:"copyright"`
	Artist    []string `json:"artist"`
	Invalid   []string `json:"invalid"`
	Lore      []string `json:"lore"`
	Meta      []string `json:"meta"`
}

func (t *TagContainer) All() []string {
	tags := t.General
	tags = append(tags, t.Species...)
	tags = append(tags, t.Character...)
	tags = append(tags, t.Copyright...)
	tags = append(tags, t.Artist...)
	tags = append(tags, t.Invalid...)
	tags = append(tags, t.Lore...)
	tags = append(tags, t.Meta...)
	return tags
}

type FlagContainer struct {
	Pending      bool `json:"pending"`
	Flagged      bool `json:"flagged"`
	NoteLocked   bool `json:"note_locked"`
	StatusLocked bool `json:"status_locked"`
	RatingLocked bool `json:"rating_locked"`
	Deleted      bool `json:"deleted"`
}

type RelationshipContainer struct {
	ParentID          interface{}   `json:"parent_id"`
	HasChildren       bool          `json:"has_children"`
	HasActiveChildren bool          `json:"has_active_children"`
	Children          []interface{} `json:"children"`
}

type PostsResponse struct {
	Posts []*Post
}
type Post struct {
	ID            int                   `json:"id"`
	CreatedAt     string                `json:"created_at"`
	UpdatedAt     string                `json:"updated_at"`
	File          File                  `json:"file"`
	Preview       File                  `json:"preview"`
	Sample        File                  `json:"sample"`
	Score         Score                 `json:"score"`
	Tags          TagContainer          `json:"tags"`
	LockedTags    []interface{}         `json:"locked_tags"` //TODO
	ChangeSeq     int                   `json:"change_seq"`
	Flags         FlagContainer         `json:"flags"`
	Rating        string                `json:"rating"`
	FavCount      int                   `json:"fav_count"`
	Sources       []string              `json:"sources"`
	Pools         []int                 `json:"pools"` // Is this type correct
	Relationships RelationshipContainer `json:"relationships"`
	ApproverID    int                   `json:"approver_id"`
	UploaderID    int                   `json:"uploader_id"`
	Description   string                `json:"description"`
	CommentCount  int                   `json:"comment_count"`
	IsFavorited   bool                  `json:"is_favorited"`
}

type TagAlias struct {
	ID             int    `json:"id"`
	AntecedentName string `json:"antecedent_name"`
	Reason         string `json:"reason"`
	CreatorID      int    `json:"creator_id"`
	CreatedAt      string `json:"created_at"`
	ForumPostID    int    `json:"forum_post_id"`
	UpdatedAt      int    `json:"updated_at"`
	ForumTopicID   int    `json:"forum_topic_id"`
	ConsequentName string `json:"consequent_name"`
	Status         string `json:"status"`
	PostCount      int    `json:"post_count"`
	ApproverID     int    `json:"approver_id"`
}

package cmd

type Options struct {
	root_MinDelay int
	root_MaxDelay int
	root_Verify   bool

	custom_Command    string
	custom_Engine     string
	custom_TargetHost string
	custom_SpiltBy    string
	custom_Depth      int
	custom_ShowHost   bool
	custom_ShowPath   bool
	custom_ShowSub    bool

	key_Domain   string
	key_Domains  string
	key_Keyword  string
	key_Keywords string
	key_Depth    int
	key_ShowHost bool
	key_ShowPath bool
	key_ShowSub  bool

	path_Domain   string
	path_Domains  string
	path_Depth    int
	path_ShowPath bool

	sub_Domain  string
	sub_Domains string
	sub_Depth   int
	sub_All     bool
	sub_ShowSub bool
	sub_ShowURL bool
}

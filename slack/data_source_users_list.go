package slack

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/slack-go/slack"
)

func dataSourceUsersList() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUsersListRead,

		Schema: map[string]*schema.Schema{
			"users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"real_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"deleted": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceUsersListRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*slack.Client)

	var users []slack.User

	us, err := client.GetUsersContext(ctx)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error fetching users list: %w", err))
	}

	var id string
	if email, ok := d.GetOk("email"); ok {
		users = filterByEmail(us, email.(string))
		id = email.(string)
	} else {
		users = us
		id = "all"
	}

	d.SetId(id)
	if err := d.Set("users", usersToMap(users)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting name: %s", err))
	}

	return diags
}

func filterByEmail(users []slack.User, email string) []slack.User {
	var res []slack.User
	for _, user := range users {
		if user.Profile.Email == email {
			res = append(res, user)
		}
	}
	return res
}

func usersToMap(users []slack.User) []map[string]interface{} {
	var res []map[string]interface{}
	for _, user := range users {
		res = append(res, userToMap(user))
	}
	return res
}

func userToMap(user slack.User) map[string]interface{} {
	res := map[string]interface{}{
		"id":        user.ID,
		"name":      user.Name,
		"email":     user.Profile.Email,
		"real_name": user.RealName,
		"deleted":   user.Deleted,
	}
	return res
}

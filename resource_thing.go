package example

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceExampleThing() *schema.Resource {
	return &schema.Resource{
		Create: resourceExampleThingCreate,
		Read:   resourceExampleThingRead,
		Update: resourceExampleThingUpdate,
		Delete: resourceExampleThingDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"thing_connection": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceExampleThingCreate(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	log.Printf("[INFO] Creating Example thing: %s", name)

	return resourceExampleThingRead(d, meta)
}

func resourceExampleThingRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceExampleThingUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Updating Example thing: %s", d.Id())
	return resourceExampleThingRead(d, meta)
}

func resourceExampleThingDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Deleting thing: %s", d.Id())
	return nil
}

func connection(d *schema.ResourceData) *ThingConnection {
	log.Printf("[INFO] setting connection: %s", d.Get("connection.0.name"))

	return &ThingConnection{
		Name: d.Get("connection.0.name").(string),
	}
}

type Thing struct {
	Name       string
	Connection *ThingConnection
}

type ThingConnection struct {
	Name string
}

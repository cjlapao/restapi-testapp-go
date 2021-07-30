package help

import (
	"runtime"
	"strings"

	"ittech24/rest/apidemo/shared"

	"github.com/fatih/color"
)

// PrintMainCommandHelper Prints specific Help
func PrintMainCommandHelper() {
	shared.LogHighlight("")
	shared.LogHighlight("Please choose a command:")
	shared.LogHighlight("")
	shared.LogHighlight("Usage:")
	shared.LogHighlight("  servicebus [command]")
	shared.LogHighlight("")
	shared.LogHighlight("Available Commands:")
	shared.LogHighlight("  topic         Service bus topic command")
	shared.LogHighlight("  queue         Service bus topic command")
}

func PrintMissingServiceBusConnectionHelper() {
	shared.LogError("Service bus connection string was not found")
	shared.Log("")
	shared.Log("Please add the SERVICEBUS_CONNECTION_STRING to your environment and try again")
	shared.Log("")
	shared.Log("Example:")
	os := runtime.GOOS
	switch strings.ToLower(os) {
	case "linux":
		shared.Log("  export SERVICEBUS_CONNECTION_STRING=\"{your connection string}\"")
	case "windows":
		shared.Log("  $env:SERVICEBUS_CONNECTION_STRING=\"{your connection string}\"")
	}
}

func PrintTopicMainCommandHelper() {
	shared.LogHighlight("")
	shared.LogHighlight("Usage:")
	shared.LogHighlight("  servicebus topic [subcommand]")
	shared.LogHighlight("")
	shared.LogHighlight("Available Sub-Commands:")
	shared.LogHighlight("  list                 Lists all Topics in a Namespace")
	shared.LogHighlight("  create               Creates a Topic in a Namespace")
	shared.LogHighlight("  delete               Deletes a Topic in a Namespace")
	shared.LogHighlight("  send                 Sends a Json Message to a specific Topic in a Namespace")
	shared.LogHighlight("  list-subscriptions   List all Subscriptions on a Topic in a Namespace")
	shared.LogHighlight("  create-subscription  Creates a Subscription on a specific Topic in a Namespace")
	shared.LogHighlight("  delete-subscription  Deletes a Subscription from a specific Topic in a Namespace")
	shared.LogHighlight("  subscribe            Subscribe to a Subscription and prints the message")
}

func PrintTopicListSubscriptionsCommandHelper() {
	shared.LogHighlight("")
	shared.LogHighlight("Usage:")
	shared.LogHighlight("  servicebus topic list-subscriptions [Options]")
	shared.LogHighlight("")
	shared.LogHighlight("Available Options:")
	shared.LogHighlight("  --name               Topic name to list subscriptions")
	shared.Log("")
	shared.Log("Example:")
	os := runtime.GOOS
	switch strings.ToLower(os) {
	case "linux":
		color.White("%v topic list-subscriptions %v", color.HiYellowString("servicebus"), color.HiBlackString("--name=example.topic"))
	case "windows":
		color.White("%v topic list-subscriptions %v", color.HiYellowString("servicebus.exe"), color.HiBlackString("--name=example.topic"))
	}
}

func PrintTopicDeleteTopicCommandHelper() {
	shared.LogHighlight("")
	shared.LogHighlight("Usage:")
	shared.LogHighlight("  servicebus topic delete [Options]")
	shared.LogHighlight("")
	shared.LogHighlight("Available Options:")
	shared.LogHighlight("  --name               Topic name to delete")
	shared.Log("")
	shared.Log("Example:")
	os := runtime.GOOS
	switch strings.ToLower(os) {
	case "linux":
		color.White("%v topic delete %v", color.HiYellowString("servicebus"), color.HiBlackString("--name=example.topic"))
	case "windows":
		color.White("%v topic delete %v", color.HiYellowString("servicebus.exe"), color.HiBlackString("--name=example.topic"))
	}
}

func PrintTopicCreateSubscriptionCommandHelper() {
	shared.LogHighlight("")
	shared.LogHighlight("Usage:")
	shared.LogHighlight("  servicebus topic create-subscription [Options]")
	shared.LogHighlight("")
	shared.LogHighlight("Available Options:")
	shared.LogHighlight("  --name                   Topic name to create subscription on")
	shared.LogHighlight("  --subscription           Subscription name to create")
	shared.LogHighlight("  --forward-to             Creates a forward to rule in the subscription")
	shared.LogHighlight("                           the format will be topic|queue:[name_of_the_target]")
	shared.LogHighlight("                           example --forward-to=topic:example.topic")
	shared.LogHighlight("  --forward-deadletter-to  Creates a forward to rule for dead letters in the subscription")
	shared.LogHighlight("                           the format will be topic|queue:[name_of_the_target]")
	shared.LogHighlight("                           example: --forward-deadletter-to=topic:example.topic")
	shared.LogHighlight("  --with-rule              Creates a sql filter/action rule for the subscription")
	shared.LogHighlight("                           the format will be [rule_name]:[sql_filter_expression]:[sql_action_expression]")
	shared.LogHighlight("                           example with only filter: --with-rule=example:2=2")
	shared.LogHighlight("                           example with filter and action: --with-rule=example:2=2:SET sys.label='example'")
	shared.Log("")
	shared.Log("Example:")
	os := runtime.GOOS
	switch strings.ToLower(os) {
	case "linux":
		color.White("%v topic create-subscription %v", color.HiYellowString("servicebus"), color.HiBlackString("--name=example.topic --subscription=example.subscription --forward-to=queue:example.queue --with-rule=example:1=1"))
	case "windows":
		color.White("%v topic create-subscription %v", color.HiYellowString("servicebus.exe"), color.HiBlackString("--name=example.topic --subscription=example.subscription --forward-to=queue:example.queue --with-rule=example:1=1"))
	}
}

func PrintTopicCreateTopicCommandHelper() {
	shared.LogHighlight("")
	shared.LogHighlight("Usage:")
	shared.LogHighlight("  servicebus topic create [Options]")
	shared.LogHighlight("")
	shared.LogHighlight("Available Options:")
	shared.LogHighlight("  --name               Topic name to create")
	shared.Log("")
	shared.Log("Example:")
	os := runtime.GOOS
	switch strings.ToLower(os) {
	case "linux":
		color.White("%v topic create %v", color.HiYellowString("servicebus"), color.HiBlackString("--name=example.topic"))
	case "windows":
		color.White("%v topic create %v", color.HiYellowString("servicebus.exe"), color.HiBlackString("--name=example.topic"))
	}
}

func PrintTopicDeleteSubscriptionCommandHelper() {
	shared.LogHighlight("")
	shared.LogHighlight("Usage:")
	shared.LogHighlight("  servicebus topic create-subscription [Options]")
	shared.LogHighlight("")
	shared.LogHighlight("Available Options:")
	shared.LogHighlight("  --name                   Topic name to delete subscription on")
	shared.LogHighlight("  --subscription           Subscription name to delete")
	shared.Log("")
	shared.Log("Example:")
	os := runtime.GOOS
	switch strings.ToLower(os) {
	case "linux":
		color.White("%v topic delete-subscription %v", color.HiYellowString("servicebus"), color.HiBlackString("--name=example.topic --subscription=example.subscription"))
	case "windows":
		color.White("%v topic delete-subscription %v", color.HiYellowString("servicebus.exe"), color.HiBlackString("--name=example.topic  --subscription=example.subscription"))
	}
}

func PrintTopicSubscribeCommandHelper() {
	shared.LogHighlight("")
	shared.LogHighlight("Usage:")
	shared.LogHighlight("  servicebus topic subscribe [options]")
	shared.LogHighlight("")
	shared.LogHighlight("Please choose a option:")
	shared.LogHighlight("")
	shared.LogHighlight("Available Options:")
	shared.LogHighlight("  %v=string         Name of the topic to listen to (mandatory)", "--topic")
	shared.LogHighlight("                         this flag can be repeated to listen to several topics")
	shared.LogHighlight("  %v=string  Name of the subscription to listen to (mandatory)", "--subscription")
	shared.LogHighlight("  %v              connects to a wiretap in the topic, if this subscription", "--wiretap")
	shared.LogHighlight("                         does not exist it will be created and deleted on exit")
	shared.LogHighlight("                         this will also override the %v flag", "--subscription")
	shared.LogHighlight("  %v                 peeks into the subscription leaving the messages there", "--peek")
	shared.LogHighlight("")
	shared.LogHighlight("example:")
	os := runtime.GOOS
	switch strings.ToLower(os) {
	case "linux":
		shared.LogHighlight("Single topic subscriber:")
		color.White("%v topic subscribe %v", color.HiYellowString("servicebus"), color.HiBlackString("--topic=example.topic --wiretap"))
		shared.LogHighlight("")
		shared.LogHighlight("Multiple topics subscriber")
		color.White("%v topic subscribe %v", color.HiYellowString("servicebus"), color.HiBlackString("--topic=example.topic --topic=example.topic2 --wiretap"))
	case "windows":
		shared.LogHighlight("Single topic subscriber:")
		color.White("%v topic subscribe %v", color.HiYellowString("servicebus.exe"), color.HiBlackString("--topic=example.topic --wiretap"))
		shared.LogHighlight("")
		shared.LogHighlight("Multiple topics subscriber")
		color.White("%v topic subscribe %v", color.HiYellowString("servicebus.exe"), color.HiBlackString("--topic=example.topic --topic=example.topic2 --wiretap"))
	}
}

func PrintTopicSendCommandHelper() {
	shared.LogHighlight("")
	shared.LogHighlight("Usage:")
	shared.LogHighlight("  servicebus topic send [options]")
	shared.LogHighlight("")
	shared.LogHighlight("Available Options:")
	shared.LogHighlight("  --topic    string     Name of the topic where to send the message")
	shared.LogHighlight("  --tenant   guid       Id of the tenant")
	shared.LogHighlight("  --body     json       Message body in json (please escape the json correctly as this is validated)")
	shared.LogHighlight("  --domain   string     Forwarding topology Message Domain")
	shared.LogHighlight("  --name     string     Forwarding topology Message Name")
	shared.LogHighlight("  --version  string     Forwarding topology Version")
	shared.LogHighlight("  --sender   string     Forwarding topology Sender")
	shared.LogHighlight("  --label    string     Message Label")
	shared.LogHighlight("  --property key:value  Add a User property to the message")
	shared.LogHighlight("                        This option can be repeated to add more than one property")
	shared.LogHighlight("                        the format will be [key]:[value]")
	shared.LogHighlight("                        example: X-Sender:example")
	shared.LogHighlight("  --default             With this flag the tool will generate a default TimeService sample")
	shared.LogHighlight("                        using the forwarding topology format")
	shared.LogHighlight("  --uno                 With this flag the tool will convert the default TimeService sample")
	shared.LogHighlight("                        message to Uno format")
	shared.LogHighlight("")
	shared.LogHighlight("example:")
	os := runtime.GOOS
	switch strings.ToLower(os) {
	case "linux":
		color.White("%v topic send %v", color.HiYellowString("servicebus"), color.HiBlackString("--topic=example.topic --body='{\\\"example\\\":\\\"document\\\"}' --domain=ExampleService --name=Example --version=\"2.1\" --sender=ExampleSender --label=ExampleLabel"))
	case "windows":
		color.White("%v topic send %v", color.HiYellowString("servicebus.exe"), color.HiBlackString("--topic=example.topic --body='{\\\"example\\\":\\\"document\\\"}' --domain=ExampleService --name=Example --version=\"2.1\" --sender=ExampleSender --label=ExampleLabel"))
	}
}

func PrintQueueMainCommandHelper() {
	shared.LogHighlight("")
	shared.LogHighlight("Usage:")
	shared.LogHighlight("  servicebus queue [subcommand]")
	shared.LogHighlight("")
	shared.LogHighlight("Available Sub-Commands:")
	shared.LogHighlight("  list                 Lists all Queues in a Namespace")
	shared.LogHighlight("  create               Creates a Queues in a Namespace")
	shared.LogHighlight("  delete               Deletes a Queues in a Namespace")
	shared.LogHighlight("  send                 Sends a Json Message to a specific Queue in a Namespace")
	shared.LogHighlight("  subscribe            Subscribe to a Queue and prints the messages")
}

func PrintQueueDeleteCommandHelper() {
	shared.LogHighlight("")
	shared.LogHighlight("Usage:")
	shared.LogHighlight("  servicebus queue delete [Options]")
	shared.LogHighlight("")
	shared.LogHighlight("Available Options:")
	shared.LogHighlight("  --name               Queue name to delete")
	shared.Log("")
	shared.Log("Example:")
	os := runtime.GOOS
	switch strings.ToLower(os) {
	case "linux":
		color.White("%v queue delete %v", color.HiYellowString("servicebus"), color.HiBlackString("--name=example.queue"))
	case "windows":
		color.White("%v queue delete %v", color.HiYellowString("servicebus.exe"), color.HiBlackString("--name=example.queue"))
	}
}

func PrintQueueCreateCommandHelper() {
	shared.LogHighlight("")
	shared.LogHighlight("Usage:")
	shared.LogHighlight("  servicebus queue create-subscription [Options]")
	shared.LogHighlight("")
	shared.LogHighlight("Available Options:")
	shared.LogHighlight("  --name                   Queue name to create subscription on")
	shared.LogHighlight("  --forward-to             Creates a forward to rule in the subscription")
	shared.LogHighlight("                           the format will be topic|queue:[name_of_the_target]")
	shared.LogHighlight("                           example --forward-to=topic:example.topic")
	shared.LogHighlight("  --forward-deadletter-to  Creates a forward to rule for dead letters in the subscription")
	shared.LogHighlight("                           the format will be topic|queue:[name_of_the_target]")
	shared.LogHighlight("                           example: --forward-deadletter-to=topic:example.topic")
	shared.Log("")
	shared.Log("Example:")
	os := runtime.GOOS
	switch strings.ToLower(os) {
	case "linux":
		color.White("%v queue create-subscription %v", color.HiYellowString("servicebus"), color.HiBlackString("--name=example.queue --forward-to=topic:example.topic --with-rule=example:1=1"))
	case "windows":
		color.White("%v queue create-subscription %v", color.HiYellowString("servicebus.exe"), color.HiBlackString("--name=example.queue --forward-to=topic:example.topic --with-rule=example:1=1"))
	}
}

func PrintQueueSendCommandHelper() {
	shared.LogHighlight("")
	shared.LogHighlight("Usage:")
	shared.LogHighlight("  servicebus queue send [options]")
	shared.LogHighlight("")
	shared.LogHighlight("Available Options:")
	shared.LogHighlight("  --queue    string     Name of the queue where to send the message")
	shared.LogHighlight("  --tenant   guid       Id of the tenant")
	shared.LogHighlight("  --body     json       Message body in json (please escape the json correctly as this is validated)")
	shared.LogHighlight("  --domain   string     Forwarding topology Message Domain")
	shared.LogHighlight("  --name     string     Forwarding topology Message Name")
	shared.LogHighlight("  --version  string     Forwarding topology Version")
	shared.LogHighlight("  --sender   string     Forwarding topology Sender")
	shared.LogHighlight("  --label    string     Message Label")
	shared.LogHighlight("  --property key:value  Add a User property to the message")
	shared.LogHighlight("                        This option can be repeated to add more than one property")
	shared.LogHighlight("                        the format will be [key]:[value]")
	shared.LogHighlight("                        example: X-Sender:example")
	shared.LogHighlight("  --default             With this flag the tool will generate a default TimeService sample")
	shared.LogHighlight("                        using the forwarding topology format")
	shared.LogHighlight("  --uno                 With this flag the tool will convert the default TimeService sample")
	shared.LogHighlight("                        message to Uno format")
	shared.LogHighlight("")
	shared.LogHighlight("example:")
	os := runtime.GOOS
	switch strings.ToLower(os) {
	case "linux":
		color.White("%v queue send %v", color.HiYellowString("servicebus"), color.HiBlackString("--queue=example.queue --body='{\\\"example\\\":\\\"document\\\"}' --domain=ExampleService --name=Example --version=\"2.1\" --sender=ExampleSender --label=ExampleLabel"))
	case "windows":
		color.White("%v queue send %v", color.HiYellowString("servicebus.exe"), color.HiBlackString("--queue=example.queue --body='{\\\"example\\\":\\\"document\\\"}' --domain=ExampleService --name=Example --version=\"2.1\" --sender=ExampleSender --label=ExampleLabel"))
	}
}

func PrintQueueSubscribeCommandHelper() {
	shared.LogHighlight("")
	shared.LogHighlight("Usage:")
	shared.LogHighlight("  servicebus queue subscribe [options]")
	shared.LogHighlight("")
	shared.LogHighlight("Please choose a option:")
	shared.LogHighlight("")
	shared.LogHighlight("Available Options:")
	shared.LogHighlight("  %v=string         Name of the queue to listen to (mandatory)", "--queue")
	shared.LogHighlight("                         this flag can be repeated to listen to several queues")
	shared.LogHighlight("  %v                 peeks into the subscription leaving the messages there", "--peek")
	shared.LogHighlight("")
	shared.LogHighlight("example:")
	os := runtime.GOOS
	switch strings.ToLower(os) {
	case "linux":
		shared.LogHighlight("Single topic subscriber:")
		color.White("%v queue subscribe %v", color.HiYellowString("servicebus"), color.HiBlackString("--queue=example.queue"))
		shared.LogHighlight("")
		shared.LogHighlight("Multiple topics subscriber")
		color.White("%v queue subscribe %v", color.HiYellowString("servicebus"), color.HiBlackString("--queue=example.queue --queue=example.queue2"))
	case "windows":
		shared.LogHighlight("Single topic subscriber:")
		color.White("%v queue subscribe %v", color.HiYellowString("servicebus.exe"), color.HiBlackString("--queue=example.queue"))
		shared.LogHighlight("")
		shared.LogHighlight("Multiple topics subscriber")
		color.White("%v queue subscribe %v", color.HiYellowString("servicebus.exe"), color.HiBlackString("--queue=example.queue --topic=example.queue2"))
	}
}

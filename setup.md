### Setting up on MACOS

This link helped me a great deal [Rabbit MQ Website](https://www.rabbitmq.com/install-homebrew.html)

1. brew update
2. brew install rabbitmq
3. set export PATH=$PATH:/usr/local/sbin in /Users/<username>/.bashrc(or .zshrc)
4. brew services start rabbitmq -> starts the server in the foreground and on system restarts


### Configuring users

One of the scripts that's installed by the brew package is rabbitmqctl, which is a tool for managing RabbitMQ nodes and used to configure all aspects of the broker.


rabbitmqctl add_user username password  (command adds user with the specified username and password)

rabbitmqctl set_user_tags miracool administrator ( giver user miracool administrator rights which is needed to access the mangement UI)

rabbitmqctl add_vhost dev-vhost (creates a v-host for development)

### Configuring dedicated vhosts

https://www.rabbitmq.com/rabbitmqctl.8.html#set_permissions tells us of how to set permissions for users in any vhost

rabbitmqctl set_permissions -p dev-vhost miracool-dev ".*" ".*" ".*" (give miracool-dev permission to do all on the dev-vhost virtual host)
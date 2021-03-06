ACM Rolebot
===========

This bot will automatically assign roles upon request for designated requestable roles.

It will also eat your vegetables upon request.

Using the bot
-------------

Type `!role list` for a list of roles.

```
me: !role list
bot: self/hello self/world
```

Type `!role get [name]` to get that role

```
me: !role get hello
bot: [thumbs up emoji]
```

Type `!role remove [name]` to remove that role

```
me: !role remove hello
bot: [thumbs up emoji]
```

Building from source
--------------------

Run `go build` or download it from [here](https://github.com/iptq/rolebot/releases/latest).

Pass your bot token in the `BOT_TOKEN` environment variable.

Run the binary.

Contact
-------

Author: Michael Zhang

License: [MIT License](https://opensource.org/licenses/MIT)

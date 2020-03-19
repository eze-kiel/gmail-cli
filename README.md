# GMail CLI

![Licence](https://img.shields.io/badge/License-GPL-brightgreen)

Send mails with your GMail account with a simple command-line interface

## Description
Typical usage :

`./main -t recipient@example.com -f yourname@gmail.com -s subject -m 'send thanks to a CLI!' -p yourGMailpassword `

Flags :
```
-f      Sender of the mail (also your username used to connect to your GMail   account)
-m      The content of your message
-p      The password of your GMail account
-s      The subject of the mail
-t      The recipient of your mail
```

Note that your password and username can be declared in environnement variables, since I'm using [namsral's flag package](github.com/namsral/flag).
## Author

Written by ezekiel.

## Copyright

License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>. This is free software: you are free to change and redistribute it.
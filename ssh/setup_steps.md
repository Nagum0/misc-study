# SSH KEY SETUP STEPS LINUX

## GENERATE SSH KEY PAIR

``` bash
ssh-keygen -t ed25519 -C "your_email@example.com"
```

- You will be prompted where to save the file (just hit enter to save it to the default place)
- You will be prompted to enter a passphrase (it can be empty if you just hit enter)

## Add Your SSH Key to the SSH Agent

- The SSH agent manages your SSH keys and remembers your passphrase so you don't have to enter it every time

``` bash
eval "$(ssh-agent -s)"
```

``` bash
ssh-add ~/.ssh/id_ed25519
```

## Access you SSH public key

``` bash
cat ~/.ssh/id_ed25519.pub
```
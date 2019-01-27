# Gimme
Use gimme to get some data out of other data sets.  Like grep but with fewer steps.

### Email
Get email from your clipboard content or from a file.

Check your clipboard for emails.  Then outputs them and copies them so they're now in your clipboard.
```
$ gimme -emails
```

Checks the file passed in with the `-f` flag for emails. Outputs the emails and copies them into the clipboard.
```
$ gimme -emails -f ./business-report.txt
```


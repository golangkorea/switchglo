# Utility for translated glossary management

# Usage

```
switchglo helps translators to manage glossary markdown. It automates followings:

1) Translators do not need to maintain the alphabetical order of terms. switchglo does it.
2) switchglo allows terms to be switched with their translated terms, vice versa.

In order to execute these automations reliably, translators should maintain following structures.

	1) Without explanation
	## Term
	translated term

	2) With explanation
	## Term
	translated term. The explanation follows.

Usage:
  switchglo [command]

Available Commands:
  sort        Sorts glossary term in alphabetical order
  switch      Switch terms with their translations

Flags:
      --config="": config file (default is $HOME/.switchglo.yaml)
      --file="glossary.md": glossary file (default is glossary.md)
  -h, --help[=false]: help for switchglo
      --out="glossary_out.md": output file (default is glossary_out.md)
  -t, --toggle[=false]: Help message for toggle

Use "switchglo [command] --help" for more information about a command.
```
# Utility for translated glossary management

## Install

```
> go get github.com/golangkorea/switchglo
```

## Usage

Once it's installed, execute in the translation project where glossary.md is present.

```
> switchglo switch
```
will swap each term with its corresponding translation, vice versa. The processed glossary will be stored in `glossary_out.md`.

Here's the help message that can be printed with `switchglo help`

```
switchglo helps translators to manage glossary markdown. It automates followings:

        1) Translators do not need to maintain the alphabetical order of terms. switchglo does it.
        2) switchglo allows terms to be switched with their translated terms, vice versa.
        3) switchglo will produce a merged glossary from existing glossary and its switched form.

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
        help        Help about any command
        merge       Merge glossary with its translated terms
        sort        Sorts glossary term in alphabetical order
        switch      Switch terms with their translations

Flags:
                --config string   config file (default is $HOME/.switchglo.yaml)
                --file string     glossary file (default is glossary.md) (default "glossary.md")
        -h, --help            help for switchglo
                --out string      output file (default is glossary_out.md) (default "glossary_out.md")
        -t, --toggle          Help message for toggle

Usage:
  switchglo [command]

Available Commands:
  help        Help about any command
  merge       Merge glossary with its translated terms
  sort        Sorts glossary term in alphabetical order
  switch      Switch terms with their translations

Flags:
      --config string   config file (default is $HOME/.switchglo.yaml)
      --file string     glossary file (default is glossary.md) (default "glossary.md")
  -h, --help            help for switchglo
      --out string      output file (default is glossary_out.md) (default "glossary_out.md")
  -t, --toggle          Help message for toggle

Use "switchglo [command] --help" for more information about a command.
```
#!/usr/bin/env bash

_console_app()
{
	local cur prev
	_get_comp_words_by_ref -n = cur prev

	commands="color clr colors \
	env ei env-info \
	example exp ex git test help"

	case "$prev" in
		example|exp)
			COMPREPLY=($(compgen -W "--id --dir --opt --names" -- "$cur"))
			return 0
			;;
	esac

	COMPREPLY=($(compgen -W "$commands" -- "$cur"))

} &&
# complete -F {auto_complete_func} {bin_filename}
complete -F _console_app cliapp

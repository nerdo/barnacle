function chpwd() {
	# save the previously set path so we can remove it from the path if necessary
	local prevNpmBinPath=$_BARNACLE_LOCAL_NPM_BIN_PATH

	_BARNACLE_LOCAL_NPM_BIN_PATH=$(barnacle -t "{{.Path}}/node_modules/.bin" node_modules/.bin package.json)

	if [[ -n $_BARNACLE_LOCAL_NPM_BIN_PATH && -d $_BARNACLE_LOCAL_NPM_BIN_PATH ]]
	then
		# https://stackoverflow.com/a/18077919
		path=($_BARNACLE_LOCAL_NPM_BIN_PATH $path)

		# https://tech.serhatteker.com/post/2019-12/remove-duplicates-in-path-zsh/
		typeset -U -g path
	else
		if [[ -n $prevNpmBinPath ]]
		then
			# https://stackoverflow.com/a/3435429
			path=("${(@)path:#$prevNpmBinPath}")
		fi
	fi
}

# prime it
chpwd

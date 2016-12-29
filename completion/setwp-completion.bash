_setwp()
{
    local cur opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    opts="--fit --stretch --center --tile --color --interval --login --wake --random --help --version"
    if [[ ${cur} == -* ]]; then
        COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    fi
}
complete -o default -F _setwp setwp

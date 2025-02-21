complete -c moneta -l amount -o amount -d 'Set amount to convert'
complete -c moneta -l from -o from -d 'Set currency to convert from (default "RUB")' --no-files -ra 'CNY EUR KZT RUB USD'
complete -c moneta -l quiet -o quiet -d 'Display less output'
complete -c moneta -s h -l help -o help -d 'Display help text'

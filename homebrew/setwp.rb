require 'formula'

class Setwp < Formula
    homepage 'https://github.com/alexandrecormier/setwp'

    ver = '0.1.1-1'
    if Hardware.is_64_bit?
        url "https://github.com/alexandrecormier/setwp/releases/download/v#{ver}/setwp-amd64-v#{ver}.tar.gz"
        sha1 'fca3b7dd0682090eced04ee277159250d851a679'
    else
        url "https://github.com/alexandrecormier/setwp/releases/download/v#{ver}/setwp-i386-v#{ver}.tar.gz"
        sha1 'd4a4f42e2eb7c827247d4946d60bc3b822b10124'
    end

    def install
        bin.install 'setwp'
        bash_completion.install 'completion/setwp-completion.bash'
        zsh_completion.install 'completion/setwp-completion.zsh'
    end

    test do
        assert_equal `#{bin}/setwp --version`.strip, "setwp version #{version}"
    end
end

require 'formula'

class Setwp < Formula
    homepage 'https://github.com/alexandrecormier/setwp'

    ver = '0.1.1-1'
    if Hardware.is_64_bit?
        url "https://github.com/alexandrecormier/setwp/releases/download/v#{ver}/setwp-amd64-v#{ver}.tar.gz"
        sha1 'a2fa531fa8a8e446ab8403e1c02123bf59aee143'
    else
        url "https://github.com/alexandrecormier/setwp/releases/download/v#{ver}/setwp-i386-v#{ver}.tar.gz"
        sha1 '1d5b9a7611bc9228ee636a67b24bffb4709c1017'
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

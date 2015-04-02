require 'formula'

class Setwp < Formula
    homepage 'https://github.com/alexandrecormier/setwp'
    version '0.1.1-1'

    if Hardware.is_64_bit?
        url "https://github.com/alexandrecormier/setwp/releases/download/v#{version}/setwp-amd64-v#{version}.tar.gz"
        sha1 '9003f374a427724782b9b49776d639b7d6987e28'
    else
        url "https://github.com/alexandrecormier/setwp/releases/download/v#{version}/setwp-i386-v#{version}.tar.gz"
        sha1 'aaa231ec22c565aef55b6313a074b4e458454527'
    end

    def install
        bin.install 'setwp'
        bash_completion.install 'completion/setwp-completion.bash'
        zsh_completion.install 'completion/setwp-completion.zsh' => '_setwp'
    end

    test do
        assert_equal `#{bin}/setwp --version`.strip, "setwp version #{version}"
    end
end

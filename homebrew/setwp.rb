require 'formula'

class Setwp < Formula
    homepage 'https://github.com/alexandrecormier/setwp'
    version '1.0.1'

    depends_on :macos => :yosemite

    url "https://github.com/alexandrecormier/setwp/releases/download/v#{version}/setwp-v#{version}.tar.gz"
    sha256 'b605eea9903a39671f61740e49188ddb9918769247251e777f272a05662bf0ab'

    def install
        bin.install 'setwp'
        bash_completion.install 'completion/setwp-completion.bash'
        zsh_completion.install 'completion/setwp-completion.zsh' => '_setwp'
    end

    def caveats ; <<-EOS.undent
        This formula will not be updated anymore and
        will eventually be removed. It is now available
        in my personal tap @ alexandrecormier/personal.
        EOS
    end

    test do
        assert_equal `#{bin}/setwp --version`.strip, "setwp version #{version}"
    end
end

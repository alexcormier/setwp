require 'formula'

class Setwp < Formula
    homepage 'https://github.com/alexandrecormier/setwp'
    version '0.1.1'

    if Hardware.is_64_bit?
        url "https://github.com/alexandrecormier/setwp/releases/download/v0.1.1/setwp-amd64-v#{version}.tar.gz"
        sha1 'a2fa531fa8a8e446ab8403e1c02123bf59aee143'
    else
        url "https://github.com/alexandrecormier/setwp/releases/download/v0.1.1/setwp-i386-v#{version}.tar.gz"
        sha1 '1d5b9a7611bc9228ee636a67b24bffb4709c1017'
    end

    def install
        bin.install 'setwp'
    end

    test do
        assert_equal `#{bin}/setwp --version`.strip, "setwp version #{version}"
    end
end

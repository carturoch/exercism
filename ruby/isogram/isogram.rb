class Isogram
  def self.isogram?(s)
    /(\w+).*\1/i !~ s
  end
end

class BookKeeping
  VERSION = 4
end
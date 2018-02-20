class Hamming
  def self.compute(str1, str2)
    raise ArgumentError unless str1.size == str2.size
    str1.chars.count {|c| c != str2.slice!(0) }
  end
end
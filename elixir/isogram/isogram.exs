defmodule Isogram do
  @doc """
  Determines if a word or sentence is an isogram
  """
  @spec isogram?(String.t) :: boolean
  def isogram?(""), do: true
  def isogram?(sentence) do
    iso = sentence
    |> String.replace(~r/\-|\s/, "")
    |> String.downcase()
    |> String.codepoints()

    uniq_count = iso 
    |> Enum.uniq() 
    |> length()

    length(iso) == uniq_count
  end

end

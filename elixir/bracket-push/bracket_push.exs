defmodule BracketPush do
  @doc """
  Checks that all the brackets and braces in the string are matched correctly, and nested correctly
  """
  @spec check_brackets(String.t()) :: boolean
  def check_brackets(str) when is_binary(str), do: check_brackets(String.codepoints(str), [])

  def check_brackets([], []), do: true
  def check_brackets([], _), do: false

  def check_brackets([h | tail], acc) when h in ~w/{ [ (/, do: check_brackets(tail, [h | acc])

  def check_brackets(["}" | tail], ["{" | rest]), do: check_brackets(tail, rest)
  def check_brackets(["]" | tail], ["[" | rest]), do: check_brackets(tail, rest)
  def check_brackets([")" | tail], ["(" | rest]), do: check_brackets(tail, rest)

  def check_brackets([h | _], _) when h in ~w/} ] )/, do: false
  def check_brackets([h | tail], acc), do: check_brackets(tail, acc)
end

defmodule Series do
  @doc """
  Finds the largest product of a given number of consecutive numbers in a given string of numbers.
  """
  @spec largest_product(String.t(), non_neg_integer) :: non_neg_integer
  def largest_product(_, 0), do: 1
  def largest_product("", _), do: raise(ArgumentError)
  def largest_product(_, size) when size < 0, do: raise(ArgumentError)

  def largest_product(number_string, size) when is_binary(number_string),
    do: largest_product(String.codepoints(number_string), size)

  def largest_product(digits, size) when size > length(digits), do: raise(ArgumentError)

  def largest_product(digits, size) do
    digits
    |> Enum.map(&String.to_integer(&1))
    |> Enum.chunk_every(size, 1)
    |> Enum.filter(&(length(&1) == size))
    |> Enum.map(fn subset -> Enum.reduce(subset, &(&1 * &2)) end)
    |> Enum.max()
  end
end

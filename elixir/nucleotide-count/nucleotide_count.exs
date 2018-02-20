defmodule NucleotideCount do
  @nucleotides [?A, ?C, ?G, ?T]

  @doc """
  Counts individual nucleotides in a NucleotideCount strand.

  ## Examples

  iex> NucleotideCount.count('AATAA', ?A)
  4

  iex> NucleotideCount.count('AATAA', ?T)
  1
  """
  @spec count([char], char) :: non_neg_integer
  def count('', _), do: 0
  def count(strand, nucleotide) do
    Enum.count(strand, &(&1 == nucleotide))
  end


  @doc """
  Returns a summary of counts by nucleotide.

  ## Examples

  iex> NucleotideCount.histogram('AATAA')
  %{?A => 4, ?T => 1, ?C => 0, ?G => 0}
  """
  @spec histogram([char]) :: map
  def histogram(''), do: %{?A => 0, ?T => 0, ?C => 0, ?G => 0}
  def histogram(strand) do
    h = histogram('')
    Enum.reduce(strand, h, fn(nucleotide, acc) -> 
      Map.put(acc, nucleotide, acc[nucleotide] + 1)
    end)
  end
end

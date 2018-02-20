defmodule ProteinTranslation do
  @stop_codons ~w/UAA UAG UGA/
  @proteins %{
    "UGU" => "Cysteine",
    "UGC" => "Cysteine",
    "UUA" => "Leucine",
    "UUG" => "Leucine",
    "AUG" => "Methionine",
    "UUU" => "Phenylalanine",
    "UUC" => "Phenylalanine",
    "UCU" => "Serine",
    "UCC" => "Serine",
    "UCA" => "Serine",
    "UCG" => "Serine",
    "UGG" => "Tryptophan",
    "UAU" => "Tyrosine",
    "UAC" => "Tyrosine"
  }

  @doc """
  Given an RNA string, return a list of proteins specified by codons, in order.
  """
  @spec of_rna(String.t()) :: {atom, list(String.t())}
  def of_rna(rna) when is_binary(rna), do: of_rna(String.codepoints(rna))

  def of_rna([a | [b | [c | _]]]) when (a <> b <> c) in @stop_codons, do: {:ok, []}
  def of_rna([]), do: {:ok, []}

  def of_rna([a | [b | [c | tail]]]) do
    case of_codon(a <> b <> c) do
      {:ok, protein} ->
        case of_rna(tail) do
          {:error, msg} ->
            {:error, msg}

          {:ok, list} ->
            {:ok, [protein | list]}
        end

      {:error, _} ->
        {:error, "invalid RNA"}
    end
  end

  @doc """
  Given a codon, return the corresponding protein

  UGU -> Cysteine
  UGC -> Cysteine
  UUA -> Leucine
  UUG -> Leucine
  AUG -> Methionine
  UUU -> Phenylalanine
  UUC -> Phenylalanine
  UCU -> Serine
  UCC -> Serine
  UCA -> Serine
  UCG -> Serine
  UGG -> Tryptophan
  UAU -> Tyrosine
  UAC -> Tyrosine
  UAA -> STOP
  UAG -> STOP
  UGA -> STOP
  """
  @spec of_codon(String.t()) :: {atom, String.t()}
  def of_codon(codon) when codon in @stop_codons, do: {:ok, "STOP"}

  def of_codon(codon) do
    case Map.fetch(@proteins, codon) do
      :error ->
        {:error, "invalid codon"}

      {:ok, protein} ->
        {:ok, protein}
    end
  end
end

defmodule Strain do
  @doc """
  Given a `list` of items and a function `fun`, return the list of items where
  `fun` returns true.

  Do not use `Enum.filter`.
  """
  @spec keep(list :: list(any), fun :: ((any) -> boolean)) :: list(any)
  def keep(list, fun), do: each(list, fun, true)

  @doc """
  Given a `list` of items and a function `fun`, return the list of items where
  `fun` returns false.

  Do not use `Enum.reject`.
  """
  @spec discard(list :: list(any), fun :: ((any) -> boolean)) :: list(any)
  def discard(list, fun), do: each(list, fun, false)

  defp each([], _, _), do: []
  defp each([h|tail], fun, flag) do
    case fun.(h) do
      ^flag -> [h|each(tail, fun, flag)]
      _ -> each(tail, fun, flag)
    end
  end
end

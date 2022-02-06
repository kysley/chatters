export const emoteSet = new Set<string>();

export const isFourPiece = (prev: string, next: string) => {
  if (!prev || !next) return;
  const prevParts = prev.split(" ");
  if (prevParts[0] === "moon21" && prevParts[1] === "moon22") {
    const nextParts = next.split(" ");
    if (nextParts[0] === "moon23") {
      return nextParts[1];
    }
  }
};

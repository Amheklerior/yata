export const play = (url: string) => {
  const audio = new Audio(url);
  audio.play().catch((e) => {
    console.error("Failed to play sound:", e);
  });
};

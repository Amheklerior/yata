import { use, useEffect, useState } from "react";
import { Toast } from "radix-ui";
import { NotificationCtx } from "../contexts/notificationCtx";

export const NotificationCenter = () => {
  const { message, clear } = use(NotificationCtx);
  const [open, setOpen] = useState(true);

  useEffect(() => {
    if (message) setOpen(true);
  }, [message]);

  const clearNotification = () => {
    setOpen(false);
    clear();
  };

  // TODO: fix enter and exit anim not working ...

  return (
    <Toast.Provider duration={3000}>
      <Toast.Root
        open={open}
        onOpenChange={clearNotification}
        className="flex items-center rounded-lg bg-red-400 px-4 py-3 shadow-lg"
      >
        <Toast.Description className="text text-red-950">
          {message}
        </Toast.Description>
      </Toast.Root>

      <Toast.Viewport className="fixed right-0 bottom-0 z-[2147483647] m-0 flex w-[390px] max-w-[100vw] list-none flex-col gap-2.5 p-[var(--viewport-padding)] outline-none [--viewport-padding:_25px]" />
    </Toast.Provider>
  );
};

import React from "react";
import { Notification } from "@mantine/core";
import { IconAlertCircle } from "@tabler/icons";

function Error({
  error,
  showError,
  setShowError,
  heading,
}: {
  error: string;
  showError: boolean;
  setShowError: React.Dispatch<React.SetStateAction<boolean>>;
  heading: string;
}) {
  return (
    <div
      hidden={!showError}
      className="absolute bottom-5 z-50 md:right-5 md:bottom-5 md:top-auto md:left-auto top-2 right-2 left-2 h-fit"
    >
      <Notification
        icon={
          <IconAlertCircle
            className="bg-red-700 text-center rounded-full h-full w-full"
            size={18}
          />
        }
        onClose={() => setShowError(false)}
        title={<h3 className="md:text-lg text-base font-bold">{heading}</h3>}
        radius="md"
        color="red"
        className="bg-red-700"
        classNames={{
          closeButton: "hover:bg-red-500",
        }}
        hidden={!showError}
      >
        <p className="md:text-base text-sm max-w-xl text-white">{error}</p>
      </Notification>
    </div>
  );
}

export default Error;

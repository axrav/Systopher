import React from "react";
import { IconCheck } from "@tabler/icons";
import { Notification } from "@mantine/core";

function Success({
  message,
  showSuccess,
  setShowSuccess,
  heading,
  loading,
}: {
  message: any;
  showSuccess: boolean;
  setShowSuccess: React.Dispatch<React.SetStateAction<boolean>>;
  heading: string;
  loading: boolean;
}) {
  return (
    <div
      hidden={!showSuccess}
      className="absolute bottom-5 z-50 md:right-5 md:bottom-5 md:top-auto md:left-auto top-2 right-2 left-2 h-fit"
    >
      <Notification
        icon={<IconCheck size={20} />}
        color="teal"
        className="bg-teal-900"
        radius="md"
        hidden={!showSuccess}
        loading={loading}
        onClose={() => setShowSuccess(false)}
        title={<h3 className="md:text-xl text-base">{heading}</h3>}
      >
        <p className="md:text-lg text-sm">{message}</p>
      </Notification>
    </div>
  );
}

export default Success;

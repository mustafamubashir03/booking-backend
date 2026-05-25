export type NotificationDto = {
  to: string;
  subject: string;
  templateId: string;
  params: Record<string, any>;
};

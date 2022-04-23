import moment from "moment";

export function parseHumanReadableDateTimeToUnixEpoch(date: string): number {
  return moment(date).unix();
}

export function parseUnixDateTime(value: number): string {
  return moment(value).format("DD/MM/YYYY, h:mm a");
}
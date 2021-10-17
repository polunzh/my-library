import dayjs from 'dayjs';

export const formatDatetime=(value)=>dayjs(value).format('YYYY-MM-DD HH:mm:ss');
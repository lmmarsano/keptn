import { Trace } from '../models/trace';


export interface EventResult {
  events: Trace[];
  totalCount: number;
}

import * as m from "minform";

export type RegisterForm = {
  email: string;
  displayName: string;
  username: string;
  password: string;
  dateOfBirth: string;
};

export const registerFormSchema = m.schema<RegisterForm>({
  email: m.string().email().required(),
  displayName: m.string().required(),
  username: m.string().required(),
  password: m.string().min(6).required(),
  dateOfBirth: m.string().required(),
});

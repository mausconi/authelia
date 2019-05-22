import Util = require("util");

export interface LdapConfiguration {
  url: string;
  base_dn: string;

  additional_users_dn?: string;
  users_filter?: string;

  additional_groups_dn?: string;
  groups_filter?: string;

  group_name_attribute?: string;
  mail_attribute?: string;

  user: string; // admin username
  password: string; // admin password
}

export function complete(configuration: LdapConfiguration): LdapConfiguration {
  const newConfiguration: LdapConfiguration = (configuration) ? JSON.parse(JSON.stringify(configuration)) : {};

  if (!newConfiguration.users_filter) {
    newConfiguration.users_filter = "cn={0}";
  }

  if (!newConfiguration.groups_filter) {
    newConfiguration.groups_filter = "member={dn}";
  }

  if (!newConfiguration.group_name_attribute) {
    newConfiguration.group_name_attribute = "cn";
  }

  if (!newConfiguration.mail_attribute) {
    newConfiguration.mail_attribute = "mail";
  }

  return newConfiguration;
}
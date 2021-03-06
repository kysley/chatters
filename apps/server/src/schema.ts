import SchemaBuilder from "@pothos/core";
import { PrismaClient } from "@prisma/client";
import PrismaPlugin from "@pothos/plugin-prisma";
import SimpleObjectsPlugin from "@pothos/plugin-simple-objects";
// This is the default location for the generator, but this can be customized as described above
// Using a type only import will help avoid issues with undeclared exports in esm mode
import type PrismaTypes from "@pothos/plugin-prisma/generated";

const prisma = new PrismaClient({});

const builder = new SchemaBuilder<{
  PrismaTypes: PrismaTypes;
}>({
  plugins: [PrismaPlugin, SimpleObjectsPlugin],
  prisma: {
    client: prisma,
  },
});

builder.prismaObject("Chatter", {
  findUnique: (chatter) => ({ id: chatter.id }),
  fields: (t) => ({
    id: t.exposeID("id"),
    username: t.exposeString("username"),
    occurances: t.relation("occurances"),
  }),
});

builder.prismaObject("Occurance", {
  findUnique: (occurance) => ({ id: occurance.id }),
  fields: (t) => ({
    id: t.exposeID("id"),
    emote: t.relation("emote"),
    uses: t.exposeInt("uses"),
    chatter: t.relation("chatter"),
  }),
});

builder.prismaObject("Emote", {
  findUnique: (emote) => ({ id: emote.id }),
  // include: {
  //   chatters: true
  // },
  fields: (t) => ({
    id: t.exposeID("id"),
    occurances: t.relation("occurances"),
    code: t.exposeString("code"),
    emoteId: t.exposeString("emoteId"),
  }),
});

const StatsType = builder.simpleObject("Stats", {
  fields: (t) => ({
    chatters: t.int(),
    occurances: t.int(),
  }),
});

builder.queryType({
  fields: (t) => ({
    chatter: t.prismaField({
      type: "Chatter",
      nullable: true,
      args: {
        username: t.arg.string({ required: true }),
      },
      resolve: async (query, root, args, _ctx, _info) => {
        return prisma.chatter.findUnique({
          where: { username: args.username },
        });
      },
    }),
    stats: t.field({
      type: StatsType,
      resolve: async (root, args) => {
        const something = await prisma.chatter.count();
        const other = await prisma.occurance.aggregate({
          _count: {
            uses: true,
          },
        });

        return {
          occurances: other._count.uses.valueOf(),
          chatters: something,
        };
      },
    }),
    uses: t.field({
      type: "Int",
      args: {
        code: t.arg.string({ required: true }),
      },
      resolve: async (root, args) => {
        const data = await prisma.occurance.aggregate({
          _sum: {
            uses: true,
          },
          where: {
            emoteCode: args.code,
          },
        });
        return data._sum.uses?.valueOf() || 0;
      },
    }),
  }),
});

export const schema = builder.toSchema({});

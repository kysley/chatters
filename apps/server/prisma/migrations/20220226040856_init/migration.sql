-- CreateTable
CREATE TABLE "Emote" (
    "id" TEXT NOT NULL,
    "emoteId" TEXT NOT NULL,
    "code" TEXT NOT NULL,

    CONSTRAINT "Emote_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Occurance" (
    "id" TEXT NOT NULL,
    "emoteCode" TEXT NOT NULL,
    "uses" INTEGER NOT NULL,
    "chatterUsername" TEXT NOT NULL,
    "updatedAt" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "Occurance_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Chatter" (
    "id" TEXT NOT NULL,
    "username" TEXT NOT NULL,

    CONSTRAINT "Chatter_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "Emote_code_key" ON "Emote"("code");

-- CreateIndex
CREATE UNIQUE INDEX "Occurance_emoteCode_chatterUsername_key" ON "Occurance"("emoteCode", "chatterUsername");

-- CreateIndex
CREATE UNIQUE INDEX "Chatter_username_key" ON "Chatter"("username");

-- AddForeignKey
ALTER TABLE "Occurance" ADD CONSTRAINT "Occurance_emoteCode_fkey" FOREIGN KEY ("emoteCode") REFERENCES "Emote"("code") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Occurance" ADD CONSTRAINT "Occurance_chatterUsername_fkey" FOREIGN KEY ("chatterUsername") REFERENCES "Chatter"("username") ON DELETE RESTRICT ON UPDATE CASCADE;

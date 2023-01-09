<template>
    <!-- comments list -->
    <div class="comments-list ">
            <div class="comments section">
                <ul>
                    <li>
                        <CustomText tag="b">_alevecchi</CustomText>
                        <span class="comment-span">It's a wonderful day</span>
                    </li>
                    <li>
                        <CustomText tag="b">m.zuckerberg</CustomText>
                        <span class="comment-span">Wait for the metaverse...</span>
                    </li>

                </ul>
            </div>
            <div class="time section">
                <CustomText size="xxsmall" class="time-ago">10 MINUTES AGO</CustomText>
            </div>
            <div class="comment section">
                <Avatar :size="20" />
                <input class="text-body" type="text" placeholder="Add a comment...">
                <a href="#" type="button">Post</a>
            </div>
        </div>
</template>
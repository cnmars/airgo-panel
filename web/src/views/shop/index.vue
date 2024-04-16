<template>
  <div class="lazy-img-container layout-pd">
    <div class="home-card-one mb15">
      <el-radio-group v-model="state.goods_type" @change="getAllEnabledGoods">
        <el-radio-button :label="constantStore.GOODS_TYPE_SUBSCRIBE" style="">
          <el-icon style="height: 100%"><Promotion /></el-icon>
          {{ $t("message.constant.GOODS_TYPE_SUBSCRIBE") }}
        </el-radio-button>
        <el-radio-button :label="constantStore.GOODS_TYPE_GENERAL">
          <el-icon><Goods /></el-icon>
          {{ $t("message.constant.GOODS_TYPE_GENERAL") }}
        </el-radio-button>
        <el-radio-button :label="constantStore.GOODS_TYPE_RECHARGE">
          <el-icon><Wallet /></el-icon>
          {{ $t("message.constant.GOODS_TYPE_RECHARGE") }}
        </el-radio-button>
      </el-radio-group>
    </div>
    <div>
      <div v-if="shopStoreData.goodsList.value.length > 0">
        <el-row :gutter="15">
          <el-col :xs="24" :sm="12" :md="8" :lg="6" :xl="4" class="mb15"
                  v-for="(v, k) in shopStoreData.goodsList.value"
                  :key="k" @click="showGoodsDetails(v)">
              <div shadow="hover" style="margin-top: 10px;border-radius:10px;background: rgba(224,224,224,0.5);padding: 20px;">
                <div >
                  <el-image :src="v.cover_image" lazy style="height: 100px;width: 100%" fit="cover">
                    <template #error>
                      <div class="image-slot">
                        <el-icon>
                          <icon-picture />
                        </el-icon>
                      </div>
                    </template>
                  </el-image>
                </div>
                <div>
                  <div style="font-weight: bolder;font-size: 20px">{{ v.subject }}</div>
                  <div style="color: #9b9da1">
                    <span>{{ $t("message.adminShop.Goods.quota") }} {{ v.quota }}</span>
                    <span style="margin-left: 20px">{{ $t("message.adminShop.Goods.stock") }} {{ v.stock }}</span>
                  </div>
                  <div class="mt15">
                      <span style="color: red">￥</span>
                      <span style="color: red;font-size: 30px">{{ v.price }}</span>
                  </div>
                </div>
              </div>

          </el-col>
        </el-row>
      </div>
      <el-empty v-else :description="$t('message.common.noData')"></el-empty>
    </div>
    <el-dialog v-model="state.isShowGoodsDetails" width="80%"
               :title="$t('message.common.details')"
               destroy-on-close>
      <el-row :gutter="50">
        <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12">
          <div style="border-radius:10px;background: rgba(224,224,224,0.29);padding: 10px">
            <div style="margin-top: 10px;text-align: center">
              <el-image :src="shopStoreData.currentGoods.value.cover_image"
                        lazy
                        style="height: 150px"
                        fit="cover"
                        :preview-src-list="[shopStoreData.currentGoods.value.cover_image]">
                <template #error>
                  <div class="image-slot">
                    <el-icon>
                      <icon-picture />
                    </el-icon>
                  </div>
                </template>
              </el-image>
            </div>
            <el-divider></el-divider>
            <div style="margin-top: 10px;">
              {{ shopStoreData.currentGoods.value.subject }}
            </div>
          </div>

          <div style="margin-top: 10px;border-radius:10px;background: rgba(224,224,224,0.29);padding: 10px">
            <div style="margin-top: 10px;margin-bottom: 10px">
              <el-tag v-if="shopStoreData.currentGoods.value.goods_type === constantStore.GOODS_TYPE_SUBSCRIBE">
                {{ $t("message.adminShop.Goods.goods_type") }}: {{ $t("message.constant.GOODS_TYPE_SUBSCRIBE") }}
              </el-tag>
              <el-tag v-if="shopStoreData.currentGoods.value.goods_type === constantStore.GOODS_TYPE_RECHARGE">
                {{ $t("message.adminShop.Goods.goods_type") }}: {{ $t("message.constant.GOODS_TYPE_RECHARGE") }}
              </el-tag>
              <el-tag v-if="shopStoreData.currentGoods.value.goods_type === constantStore.GOODS_TYPE_GENERAL">
                {{ $t("message.adminShop.Goods.goods_type") }}: {{ $t("message.constant.GOODS_TYPE_GENERAL") }}
              </el-tag>

              <el-tag v-if="shopStoreData.currentGoods.value.deliver_type === constantStore.DELIVER_TYPE_NONE">
                {{ $t("message.adminShop.Goods.deliver_type") }}: {{ $t("message.constant.DELIVER_TYPE_NONE") }}
              </el-tag>
              <el-tag v-if="shopStoreData.currentGoods.value.deliver_type === constantStore.DELIVER_TYPE_AUTO">
                {{ $t("message.adminShop.Goods.deliver_type") }}: {{ $t("message.constant.DELIVER_TYPE_AUTO") }}
              </el-tag>
              <el-tag v-if="shopStoreData.currentGoods.value.deliver_type === constantStore.DELIVER_TYPE_MANUAL">
                {{ $t("message.adminShop.Goods.deliver_type") }}: {{ $t("message.constant.DELIVER_TYPE_MANUAL") }}
              </el-tag>
            </div>
            <div style="margin-top: 10px;margin-bottom: 10px">
              <el-tag type="warning">{{ $t("message.adminShop.Goods.quota") }}：{{ shopStoreData.currentGoods.value.quota
                }} / {{ $t("message.adminShop.Goods.stock") }}：{{ shopStoreData.currentGoods.value.stock }}
              </el-tag>
            </div>
            <el-descriptions
              :column="1"
              border
              size="small"
              direction="horizontal"
            >
              <div v-if="shopStoreData.currentGoods.value.goods_type === constantStore.GOODS_TYPE_SUBSCRIBE">
                <el-descriptions-item :label="$t('message.adminShop.Goods.total_bandwidth')">
                  {{ shopStoreData.currentGoods.value.total_bandwidth }}GB
                </el-descriptions-item>
                <el-descriptions-item :label="$t('message.adminShop.Goods.node_connector')">
                  {{ shopStoreData.currentGoods.value.node_connector }}
                </el-descriptions-item>
                <el-descriptions-item :label="$t('message.adminShop.Goods.node_speed_limit')">
                  {{ shopStoreData.currentGoods.value.node_speed_limit }}
                </el-descriptions-item>
              </div>
              <div v-if="shopStoreData.currentGoods.value.goods_type === constantStore.GOODS_TYPE_RECHARGE">
                <el-descriptions-item :label="$t('message.adminShop.Goods.recharge_amount')">
                  {{ shopStoreData.currentGoods.value.recharge_amount }}
                </el-descriptions-item>
              </div>
            </el-descriptions>
            <div style="margin-top: 10px">
                <span>
                  <span style="color: red;">￥</span>
                  <span style="color: red;font-size: 30px;">{{ shopStoreData.currentGoods.value.price }}</span>
                  <span
                    v-if="shopStoreData.currentGoods.value.goods_type === constantStore.GOODS_TYPE_SUBSCRIBE"> / {{ $t("message.common.month") }}</span>
                </span>
            </div>
          </div>
        </el-col>

        <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12">
          <div style="margin-top: 10px;border-radius:10px;background:rgba(224,224,224,0.29)">
            <div style="margin-top: 10px;padding: 10px" v-html="shopStoreData.currentGoods.value.des"></div>
          </div>
        </el-col>
      </el-row>
      <template #footer>
				<span>
					<el-button @click="state.isShowGoodsDetails = false"
                     size="default">{{ $t("message.common.button_cancel") }}
          </el-button>
					<el-button type="primary"
                     :disabled="shopStoreData.currentGoods.value.stock <= 0"
                     @click="openPurchase(shopStoreData.currentGoods.value)"
                     size="default">{{ $t("message.adminShop.purchase") }}
          </el-button>
				</span>
      </template>
    </el-dialog>
    <Purchase ref="PurchaseRef"></Purchase>
  </div>
</template>

<script setup lang="ts" name="pagesLazyImg">
import { reactive, onMounted, ref, defineAsyncComponent } from "vue";
import { useRouter } from "vue-router";
// import other from '/@/utils/other';

import { useShopStore } from "/@/stores/user_logic/shopStore";
import { storeToRefs } from "pinia";
import { useConstantStore } from "/@/stores/constantStore";

const shopStore = useShopStore();
const shopStoreData = storeToRefs(shopStore);
const constantStore = useConstantStore();
const Purchase = defineAsyncComponent(() => import("/@/views/shop/purchase.vue"));
const PurchaseRef = ref();

// 定义变量内容
const router = useRouter();
const state = reactive({
  isShowGoodsDetails: false,
  isShowLoading: false,
  goods_type: constantStore.GOODS_TYPE_SUBSCRIBE
});

const showGoodsDetails = (v: Goods) => {
  shopStoreData.currentGoods.value = v;
  state.isShowGoodsDetails = true;
};
const openPurchase = () => {
  state.isShowGoodsDetails = false;
  state.isShowLoading = true;
  shopStoreData.currentOrder.value.duration = 1; // 默认订购时长
  shopStoreData.currentOrder.value.order_type = constantStore.ORDER_TYPE_NEW; //订单类型：新购入
  shopStoreData.currentOrder.value.goods_id = shopStoreData.currentGoods.value.id; //订购商品ID

  PurchaseRef.value.openDialog(constantStore.ORDER_TYPE_NEW);
};
const getAllEnabledGoods = () => {
  shopStore.getAllEnabledGoods({ goods_type: state.goods_type });
};
// 页面加载时
onMounted(() => {
  getAllEnabledGoods();
});
</script>

<style scoped lang="scss">

</style>

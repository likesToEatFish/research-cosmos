package keeper

import (
	"context"
	"fmt"
	"strconv"
	"time"

	errorsmod "cosmossdk.io/errors"
	math "cosmossdk.io/math"
	prefix "cosmossdk.io/store/prefix"
	runtime "github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
	"github.com/onomyprotocol/reserve/x/oracle/types"
)

// SetBandParams sets the Band params in the state
func (k Keeper) SetBandParams(ctx context.Context, bandParams types.BandParams) error {
	bz := k.cdc.MustMarshal(&bandParams)
	store := k.storeService.OpenKVStore(ctx)
	return store.Set(types.BandParamsKey, bz)
}

// GetBandParams gets the Band params stored in the state
func (k Keeper) GetBandParams(ctx context.Context) types.BandParams {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get(types.BandParamsKey)

	if err != nil {
		return types.DefaultGenesis().BandParams
	}

	if bz == nil {
		return types.DefaultGenesis().BandParams
	}

	var bandParams types.BandParams
	k.cdc.MustUnmarshal(bz, &bandParams)
	return bandParams
}

// SetBandOracleRequestParams sets the Band Oracle request params in the state
func (k Keeper) SetBandOracleRequestParams(ctx context.Context, bandOracleRequestParams types.BandOracleRequestParams) error {
	bz := k.cdc.MustMarshal(&bandOracleRequestParams)
	store := k.storeService.OpenKVStore(ctx)
	return store.Set(types.BandOracleRequestParamsKey, bz)
}

// GetBandParams gets the Band params stored in the state
func (k Keeper) GetBandOracleRequestParams(ctx context.Context) types.BandOracleRequestParams {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get(types.BandOracleRequestParamsKey)

	if err != nil {
		return types.DefaultGenesis().BandOracleRequestParams
	}

	if bz == nil {
		return types.DefaultGenesis().BandOracleRequestParams
	}

	var bandOracleRequestParams types.BandOracleRequestParams
	k.cdc.MustUnmarshal(bz, &bandOracleRequestParams)
	return bandOracleRequestParams
}

// SetBandCallData sets the Band IBC oracle request call data
func (k Keeper) SetBandCallDataRecord(ctx context.Context, record *types.CalldataRecord) error {
	bz := k.cdc.MustMarshal(record)
	store := k.storeService.OpenKVStore(ctx)
	return store.Set(types.GetBandCallDataRecordKey(record.ClientId), bz)
}

// DeleteBandCallDataRecord deletes the Band IBC oracle request call data
func (k Keeper) DeleteBandCallDataRecord(ctx context.Context, clientID uint64) error {
	store := k.storeService.OpenKVStore(ctx)
	return store.Delete(types.GetBandCallDataRecordKey(clientID))
}

// GetAllBandCalldataRecords gets all Band oracle request CallData for each clientID
func (k Keeper) GetAllBandCalldataRecords(ctx context.Context) []*types.CalldataRecord {
	calldataRecords := make([]*types.CalldataRecord, 0)
	kvStore := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bandCalldataStore := prefix.NewStore(kvStore, types.BandCallDataRecordKey)

	iterator := bandCalldataStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var calldataRecord types.CalldataRecord
		k.cdc.MustUnmarshal(iterator.Value(), &calldataRecord)
		calldataRecords = append(calldataRecords, &calldataRecord)
	}

	return calldataRecords
}

// GetBandCallDataRecord gets the Band oracle request CallDataRecord for a given clientID
func (k Keeper) GetBandCallDataRecord(ctx context.Context, clientID uint64) *types.CalldataRecord {
	var callDataRecord types.CalldataRecord
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get(types.GetBandCallDataRecordKey(clientID))
	if err != nil {
		return nil
	}
	if bz == nil {
		return nil
	}
	k.cdc.MustUnmarshal(bz, &callDataRecord)
	return &callDataRecord
}

// GetBandLatestClientID returns the latest clientID of Band oracle request packet data.
func (k Keeper) GetBandLatestClientID(ctx context.Context) uint64 {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get(types.LatestClientIDKey)
	if err != nil {
		return 0
	}
	if bz == nil {
		return 0
	}
	clientID := sdk.BigEndianToUint64(bz)
	return clientID
}

// SetBandLatestClientID sets the latest clientID of Band oracle request packet data.
func (k Keeper) SetBandLatestClientID(ctx context.Context, clientID uint64) error {
	store := k.storeService.OpenKVStore(ctx)
	return store.Set(types.LatestClientIDKey, sdk.Uint64ToBigEndian(clientID))
}

// GetBandLatestRequestID returns the latest requestID of Band oracle request types.
func (k Keeper) GetBandLatestRequestID(ctx context.Context) uint64 {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get(types.LatestRequestIDKey)
	if err != nil {
		// TODO: should we return 0 here?
		return 0
	}
	if bz == nil {
		return 0
	}
	requestID := sdk.BigEndianToUint64(bz)
	return requestID
}

// SetBandLatestRequestID sets the latest requestID of Band oracle request types.
func (k Keeper) SetBandLatestRequestID(ctx context.Context, requestID uint64) error {
	store := k.storeService.OpenKVStore(ctx)
	return store.Set(types.LatestRequestIDKey, sdk.Uint64ToBigEndian(requestID))
}

// SetBandOracleRequest sets the Band oracle request data
func (k Keeper) SetBandOracleRequest(ctx context.Context, req types.BandOracleRequest) error {
	bz := k.cdc.MustMarshal(&req)
	store := k.storeService.OpenKVStore(ctx)
	return store.Set(types.GetBandOracleRequestIDKey(req.RequestId), bz)
}

// GetBandOracleRequest gets the Band oracle request data
func (k Keeper) GetBandOracleRequest(ctx context.Context, requestID uint64) *types.BandOracleRequest {
	var bandOracleRequest types.BandOracleRequest
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get(types.GetBandOracleRequestIDKey(requestID))
	if err != nil {
		return nil
	}
	if bz == nil {
		return nil
	}

	k.cdc.MustUnmarshal(bz, &bandOracleRequest)
	return &bandOracleRequest
}

// DeleteBandOracleRequest deletes the Band oracle request call data
func (k Keeper) DeleteBandOracleRequest(ctx context.Context, requestID uint64) error {
	store := k.storeService.OpenKVStore(ctx)
	return store.Delete(types.GetBandOracleRequestIDKey(requestID))
}

// GetAllBandOracleRequests gets all Band oracle requests for each requestID
func (k Keeper) GetAllBandOracleRequests(ctx context.Context) []*types.BandOracleRequest {
	bandOracleRequests := make([]*types.BandOracleRequest, 0)
	kvStore := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bandOracleRequestStore := prefix.NewStore(kvStore, types.BandOracleRequestIDKey)

	iterator := bandOracleRequestStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var bandOracleRequest types.BandOracleRequest
		k.cdc.MustUnmarshal(iterator.Value(), &bandOracleRequest)
		bandOracleRequests = append(bandOracleRequests, &bandOracleRequest)
	}

	return bandOracleRequests
}

// GetBandPriceState reads the stored band ibc price state.
func (k *Keeper) GetBandPriceState(ctx context.Context, symbol string) *types.BandPriceState {
	var priceState types.BandPriceState
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get(types.GetBandPriceStoreKey(symbol))
	if err != nil {
		return nil
	}
	if bz == nil {
		return nil
	}

	k.cdc.MustUnmarshal(bz, &priceState)
	return &priceState
}

// SetBandPriceState sets the band ibc price state.
func (k *Keeper) SetBandPriceState(ctx context.Context, symbol string, priceState *types.BandPriceState) error {
	bz := k.cdc.MustMarshal(priceState)
	store := k.storeService.OpenKVStore(ctx)
	return store.Set(types.GetBandPriceStoreKey(symbol), bz)
}

// GetAllBandPriceStates reads all stored band price states.
func (k *Keeper) GetAllBandPriceStates(ctx context.Context) []*types.BandPriceState {
	priceStates := make([]*types.BandPriceState, 0)
	kvStore := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bandPriceStore := prefix.NewStore(kvStore, types.BandPriceKey)

	iterator := bandPriceStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var bandPriceState types.BandPriceState
		k.cdc.MustUnmarshal(iterator.Value(), &bandPriceState)
		priceStates = append(priceStates, &bandPriceState)
	}

	return priceStates
}

// AddNewSymbolToBandOracleRequest adds a new symbol to the bandOracle request
func (k Keeper) AddNewSymbolToBandOracleRequest(ctx context.Context, symbol string, oracleScriptId int64) error {
	allBandOracleRequests := k.GetAllBandOracleRequests(ctx)
	// check if new symbol's oracle script id is existing
	for _, req := range allBandOracleRequests {
		if req.OracleScriptId == oracleScriptId {
			req.Symbols = append(req.Symbols, symbol)
			k.SetBandOracleRequest(ctx, *req)
			return nil
		}
	}

	bandOracleRequestParams := k.GetBandOracleRequestParams(ctx)
	requestID := k.GetBandLatestRequestID(ctx) + 1
	newBandOracleRequest := types.BandOracleRequest{
		RequestId:      requestID,
		OracleScriptId: oracleScriptId,
		Symbols:        []string{symbol},
		AskCount:       bandOracleRequestParams.AskCount,
		MinCount:       bandOracleRequestParams.MinCount,
		FeeLimit:       bandOracleRequestParams.FeeLimit,
		PrepareGas:     bandOracleRequestParams.PrepareGas,
		ExecuteGas:     bandOracleRequestParams.ExecuteGas,
		MinSourceCount: bandOracleRequestParams.MinSourceCount,
	}

	k.SetBandOracleRequest(ctx, newBandOracleRequest)

	k.SetBandLatestRequestID(ctx, requestID)
	return nil
}

// GetPrice fetches band ibc prices for a given pair in math.LegacyDec
func (k Keeper) GetPrice(ctx context.Context, base, quote string) *math.LegacyDec {
	// query ref by using GetBandPriceState
	basePriceState := k.GetBandPriceState(ctx, base)
	if basePriceState == nil || basePriceState.Rate.IsZero() {
		k.Logger(ctx).Info("Can not get price state of base denom %s: price state is nil or rate is zero", base)
		return nil
	}

	if quote == types.QuoteUSD {
		return &basePriceState.PriceState.Price
	}

	quotePriceState := k.GetBandPriceState(ctx, quote)
	if quotePriceState == nil || quotePriceState.Rate.IsZero() {
		k.Logger(ctx).Info("Can not get price state of quote denom %s: price state is nil or rate is zero", quote)
		return nil
	}

	baseRate := basePriceState.Rate.ToLegacyDec()
	quoteRate := quotePriceState.Rate.ToLegacyDec()

	if baseRate.IsNil() || quoteRate.IsNil() || !baseRate.IsPositive() || !quoteRate.IsPositive() {
		return nil
	}

	price := baseRate.Quo(quoteRate)
	return &price
}

// RequestBandOraclePrices creates and sends an IBC packet to fetch band oracle price feed data through IBC.
func (k *Keeper) RequestBandOraclePrices(
	ctx context.Context,
	req *types.BandOracleRequest,
) (err error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	bandParams := k.GetBandParams(ctx)
	sourcePortID := bandParams.IbcPortId
	sourceChannel := bandParams.IbcSourceChannel

	calldata := req.GetCalldata(types.IsLegacySchemeOracleScript(req.OracleScriptId, bandParams))

	sourceChannelEnd, found := k.channelKeeper.GetChannel(sdkCtx, sourcePortID, sourceChannel)
	if !found {
		return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unknown channel %s port %s", sourceChannel, sourcePortID)
	}

	// retrieve the dynamic capability for this channel
	channelCap, ok := k.scopedKeeper.GetCapability(sdkCtx, host.ChannelCapabilityPath(sourcePortID, sourceChannel))
	if !ok {
		return errorsmod.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	destinationPort := sourceChannelEnd.Counterparty.PortId
	destinationChannel := sourceChannelEnd.Counterparty.ChannelId
	sequence, found := k.channelKeeper.GetNextSequenceSend(sdkCtx, sourcePortID, sourceChannel)

	if !found {
		return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unknown sequence number for channel %s port %s", sourceChannel, sourcePortID)
	}

	clientID := k.GetBandLatestClientID(ctx) + 1
	packetData := types.NewOracleRequestPacketData(strconv.Itoa(int(clientID)), calldata, req)

	// Creating custom oracle packet data
	packet := channeltypes.NewPacket(
		packetData.GetBytes(),
		sequence,
		sourcePortID,
		sourceChannel,
		destinationPort,
		destinationChannel,
		clienttypes.NewHeight(0, 0),
		uint64(sdkCtx.BlockTime().UnixNano()+int64(20*time.Minute)), // Arbitrarily high timeout for now
	)

	// Send packet to IBC, authenticating with channelCap
	_, err = k.channelKeeper.SendPacket(
		sdkCtx,
		channelCap,
		packet.SourcePort,
		packet.SourceChannel,
		packet.TimeoutHeight,
		packet.TimeoutTimestamp,
		packet.Data,
	)
	if err != nil {
		return err
	}

	// Persist the sequence number and OracleRequest CallData. CallData contains list of symbols.
	// This is used to map the prices/rates with the symbols upon receiving oracle response from Band IBC.
	k.SetBandCallDataRecord(ctx, &types.CalldataRecord{
		ClientId: clientID,
		Calldata: calldata,
	})

	k.SetBandLatestClientID(ctx, clientID)

	return
}

func (k *Keeper) ProcessBandOraclePrices(
	ctx sdk.Context,
	relayer sdk.Address,
	packet types.OracleResponsePacketData,
) error {
	clientID, err := strconv.Atoi(packet.ClientID)
	if err != nil {
		return fmt.Errorf("failed to parse client ID: %w", err)
	}

	callRecord := k.GetBandCallDataRecord(ctx, uint64(clientID))
	if callRecord == nil {
		// TODO: should this be an error?
		return nil
	}

	input, err := types.DecodeOracleInput(callRecord.Calldata)
	if err != nil {
		return err
	}

	output, err := types.DecodeOracleOutput(packet.Result)
	if err != nil {
		return err
	}

	k.updateBandPriceStates(ctx, input, output, packet, relayer, clientID)

	// Delete the calldata corresponding to the sequence number
	k.DeleteBandCallDataRecord(ctx, uint64(clientID))

	return nil
}

func (k *Keeper) updateBandPriceStates(
	ctx sdk.Context,
	input types.OracleInput,
	output types.OracleOutput,
	packet types.OracleResponsePacketData,
	relayer sdk.Address,
	clientID int,
) {
	var (
		inputSymbols = input.PriceSymbols()
		requestID    = packet.RequestID
		resolveTime  = uint64(packet.ResolveTime)
		symbols      = make([]string, 0, len(inputSymbols))
		prices       = make([]math.LegacyDec, 0, len(inputSymbols))
	)

	// loop SetBandPriceState for all symbols
	for idx, symbol := range inputSymbols {
		if !output.Valid(idx) {
			//	failed response for given symbol, skip it
			continue
		}

		var (
			rate       = output.Rate(idx)
			multiplier = input.PriceMultiplier()
			price      = math.LegacyNewDec(int64(rate)).Quo(math.LegacyNewDec(int64(multiplier)))
		)
		if price.IsZero() {
			continue
		}

		bandPriceState := k.GetBandPriceState(ctx, symbol)

		// don't update band prices with an older price
		if bandPriceState != nil && bandPriceState.ResolveTime > resolveTime {
			continue
		}

		// skip price update if the price changes beyond 100x or less than 1% of the last price
		if bandPriceState != nil && types.CheckPriceFeedThreshold(bandPriceState.PriceState.Price, price) {
			continue
		}

		blockTime := ctx.BlockTime().Unix()
		if bandPriceState == nil {
			bandPriceState = &types.BandPriceState{
				Symbol:      symbol,
				Rate:        math.NewInt(int64(rate)),
				ResolveTime: resolveTime,
				Request_ID:  requestID,
				PriceState:  *types.NewPriceState(price, blockTime),
			}
		} else {
			bandPriceState.Rate = math.NewInt(int64(rate))
			bandPriceState.ResolveTime = resolveTime
			bandPriceState.Request_ID = requestID
			bandPriceState.PriceState.UpdatePrice(price, blockTime)
		}

		err := k.SetBandPriceState(ctx, symbol, bandPriceState)
		if err != nil {
			k.Logger(ctx).Info("Can not set band price state for symbol %v", symbol)
		}

		symbols = append(symbols, symbol)
		prices = append(prices, price)
	}

	if len(symbols) == 0 {
		return
	}

	// emit SetBandPriceEvent event
	// nolint:errcheck //ignored on purpose
	ctx.EventManager().EmitTypedEvent(&types.SetBandPriceEvent{
		Relayer:     relayer.String(),
		Symbols:     symbols,
		Prices:      prices,
		ResolveTime: uint64(packet.ResolveTime),
		RequestId:   packet.RequestID,
		ClientId:    int64(clientID),
	})
}

func (k *Keeper) CleanUpStaleBandCalldataRecords(ctx context.Context) {
	var (
		latestClientID         = k.GetBandLatestClientID(ctx)
		earliestToKeepClientID = latestClientID - 1000 // todo: default max records to keep (1000)
	)

	if earliestToKeepClientID > latestClientID {
		// underflow
		return
	}

	for _, id := range k.getPreviousRecordIDs(ctx, earliestToKeepClientID) {
		k.DeleteBandCallDataRecord(ctx, id)
	}
}

func (k *Keeper) getPreviousRecordIDs(ctx context.Context, clientID uint64) []uint64 {
	kvStore := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bandCalldataStore := prefix.NewStore(kvStore, types.BandCallDataRecordKey)
	iter := bandCalldataStore.Iterator(nil, sdk.Uint64ToBigEndian(clientID))
	defer iter.Close()

	staleIDs := make([]uint64, 0)
	for ; iter.Valid(); iter.Next() {
		var record types.CalldataRecord
		k.cdc.MustUnmarshal(iter.Value(), &record)

		staleIDs = append(staleIDs, record.ClientId)
	}

	return staleIDs
}

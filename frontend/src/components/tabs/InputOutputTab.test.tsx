/*
 * Copyright 2021 The Kubeflow Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { render, screen, waitFor } from '@testing-library/react';
import { Struct } from 'google-protobuf/google/protobuf/struct_pb';
import React from 'react';
import { Apis } from 'src/lib/Apis';
import { Api } from 'src/mlmd/library';
import { testBestPractices } from 'src/TestUtils';
import { CommonTestWrapper } from 'src/TestWrapper';
import {
  Artifact,
  Event,
  Execution,
  GetArtifactsByIDResponse,
  GetEventsByExecutionIDsResponse,
  Value,
} from 'src/third_party/mlmd';
import InputOutputTab from './InputOutputTab';

const executionName = 'fake-execution';
const artifactId = 100;
const artifactUri = 'gs://bucket/test';
const artifactUriView = 'gs://bucket/test';
const inputArtifactName = 'input_artifact';
const outputArtifactName = 'output_artifact';
const namespace = 'namespace';

testBestPractices();
describe('InoutOutputTab', () => {
  it('shows execution title', () => {
    jest
      .spyOn(Api.getInstance().metadataStoreService, 'getEventsByExecutionIDs')
      .mockResolvedValue(new GetEventsByExecutionIDsResponse());
    jest
      .spyOn(Api.getInstance().metadataStoreService, 'getArtifactsByID')
      .mockReturnValue(new GetArtifactsByIDResponse());

    render(
      <CommonTestWrapper>
        <InputOutputTab execution={buildBasicExecution()} namespace={namespace}></InputOutputTab>
      </CommonTestWrapper>,
    );
    screen.getByText(executionName, { selector: 'a', exact: false });
  });

  it("doesn't show Input/Output artifacts and parameters if no exists", async () => {
    jest
      .spyOn(Api.getInstance().metadataStoreService, 'getEventsByExecutionIDs')
      .mockResolvedValue(new GetEventsByExecutionIDsResponse());
    jest
      .spyOn(Api.getInstance().metadataStoreService, 'getArtifactsByID')
      .mockReturnValue(new GetArtifactsByIDResponse());

    render(
      <CommonTestWrapper>
        <InputOutputTab execution={buildBasicExecution()} namespace={namespace}></InputOutputTab>
      </CommonTestWrapper>,
    );
    await waitFor(() => screen.queryAllByText('Input Parameters').length == 0);
    await waitFor(() => screen.queryAllByText('Input Artifacts').length == 0);
    await waitFor(() => screen.queryAllByText('Output Parameters').length == 0);
    await waitFor(() => screen.queryAllByText('Output Artifacts').length == 0);
    await waitFor(() => screen.getByText('There is no input/output parameter or artifact.'));
  });

  it('shows Input parameters with various types', async () => {
    jest
      .spyOn(Api.getInstance().metadataStoreService, 'getEventsByExecutionIDs')
      .mockResolvedValue(new GetEventsByExecutionIDsResponse());
    jest
      .spyOn(Api.getInstance().metadataStoreService, 'getArtifactsByID')
      .mockReturnValue(new GetArtifactsByIDResponse());

    const execution = buildBasicExecution();
    execution
      .getCustomPropertiesMap()
      .set('thisKeyIsNotInput', new Value().setStringValue("value shouldn't show"));
    execution.getCustomPropertiesMap().set(
      'inputs',
      new Value().setStructValue(
        Struct.fromJavaScript({
          'stringKey: example string': 0,
          'intKey: 123': 1,
          'doubleKey: 1.23': 2,
          'listKey: [1,2,3]': 3,
          'structKey: {a:1, b:2}': 4,
        }),
      ),
    );

    render(
      <CommonTestWrapper>
        <InputOutputTab execution={execution} namespace={namespace}></InputOutputTab>
      </CommonTestWrapper>,
    );

    screen.getByText('stringKey: example string');
    screen.getByText('intKey: 123');
    screen.getByText('doubleKey: 1.23');
    screen.getByText('listKey: [1,2,3]');
    screen.getByText('structKey: {a:1, b:2}');
    expect(screen.queryByText('thisKeyIsNotInput')).toBeNull();
  });

  it('shows Output parameters with various types', async () => {
    jest
      .spyOn(Api.getInstance().metadataStoreService, 'getEventsByExecutionIDs')
      .mockResolvedValue(new GetEventsByExecutionIDsResponse());
    jest
      .spyOn(Api.getInstance().metadataStoreService, 'getArtifactsByID')
      .mockReturnValue(new GetArtifactsByIDResponse());

    const execution = buildBasicExecution();
    execution
      .getCustomPropertiesMap()
      .set('thisKeyIsNotOutput', new Value().setStringValue("value shouldn't show"));
    execution.getCustomPropertiesMap().set(
      'outputs',
      new Value().setStructValue(
        Struct.fromJavaScript({
          'stringKey: example string': 0,
          'intKey: 123': 1,
          'doubleKey: 1.23': 2,
          'listKey: [1,2,3]': 3,
          'structKey: {a:1, b:2}': 4,
        }),
      ),
    );

    render(
      <CommonTestWrapper>
        <InputOutputTab execution={execution} namespace={namespace}></InputOutputTab>
      </CommonTestWrapper>,
    );

    screen.getByText('stringKey: example string');
    screen.getByText('intKey: 123');
    screen.getByText('doubleKey: 1.23');
    screen.getByText('listKey: [1,2,3]');
    screen.getByText('structKey: {a:1, b:2}');
    expect(screen.queryByText('thisKeyIsNotOutput')).toBeNull();
  });

  it('shows Input artifacts', async () => {
    jest.spyOn(Apis, 'readFile').mockResolvedValue('artifact preview');
    const getEventResponse = new GetEventsByExecutionIDsResponse();
    getEventResponse.getEventsList().push(buildInputEvent());
    jest
      .spyOn(Api.getInstance().metadataStoreService, 'getEventsByExecutionIDs')
      .mockResolvedValueOnce(getEventResponse);
    const getArtifactsResponse = new GetArtifactsByIDResponse();
    getArtifactsResponse.getArtifactsList().push(buildArtifact());
    jest
      .spyOn(Api.getInstance().metadataStoreService, 'getArtifactsByID')
      .mockReturnValueOnce(getArtifactsResponse);

    render(
      <CommonTestWrapper>
        <InputOutputTab execution={buildBasicExecution()} namespace={namespace}></InputOutputTab>
      </CommonTestWrapper>,
    );

    await waitFor(() => screen.getByText(artifactUriView));
    await waitFor(() => screen.getByText(inputArtifactName));
  });

  it('shows Output artifacts', async () => {
    jest.spyOn(Apis, 'readFile').mockResolvedValue('artifact preview');
    const getEventResponse = new GetEventsByExecutionIDsResponse();
    getEventResponse.getEventsList().push(buildOutputEvent());
    jest
      .spyOn(Api.getInstance().metadataStoreService, 'getEventsByExecutionIDs')
      .mockResolvedValueOnce(getEventResponse);
    const getArtifactsResponse = new GetArtifactsByIDResponse();
    getArtifactsResponse.getArtifactsList().push(buildArtifact());
    jest
      .spyOn(Api.getInstance().metadataStoreService, 'getArtifactsByID')
      .mockReturnValueOnce(getArtifactsResponse);

    render(
      <CommonTestWrapper>
        <InputOutputTab execution={buildBasicExecution()} namespace={namespace}></InputOutputTab>
      </CommonTestWrapper>,
    );

    await waitFor(() => screen.getByText(artifactUriView));
    await waitFor(() => screen.getByText(outputArtifactName));
  });
});

function buildBasicExecution() {
  const execution = new Execution();
  const executionId = 123;

  execution.setId(executionId);
  execution.getCustomPropertiesMap().set('task_name', new Value().setStringValue(executionName));

  return execution;
}

function buildArtifact() {
  const artifact = new Artifact();
  artifact.getCustomPropertiesMap();
  artifact.setUri(artifactUri);
  artifact.setId(artifactId);
  return artifact;
}

function buildInputEvent() {
  const event = new Event();
  const path = new Event.Path();
  path.getStepsList().push(new Event.Path.Step().setKey(inputArtifactName));
  event
    .setType(Event.Type.INPUT)
    .setArtifactId(artifactId)
    .setPath(path);
  return event;
}

function buildOutputEvent() {
  const event = new Event();
  const path = new Event.Path();
  path.getStepsList().push(new Event.Path.Step().setKey(outputArtifactName));
  event
    .setType(Event.Type.OUTPUT)
    .setArtifactId(artifactId)
    .setPath(path);
  return event;
}
